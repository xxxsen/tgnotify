//Package dao db处理类
package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"tgnotify/models"
)

type FileStorage struct {
	file          string
	indexByChatid map[uint64]*models.UserInfo
	lck           sync.RWMutex
}

var defaultFs *FileStorage

func NewFileStorage(file string) (*FileStorage, error) {
	fs := &FileStorage{
		file:          file,
		indexByChatid: make(map[uint64]*models.UserInfo),
	}
	if err := fs.initFs(); err != nil {
		return nil, err
	}
	return fs, nil

}

func Init(target string) error {
	fs, err := NewFileStorage(target)
	defaultFs = fs
	return err
}

func GetFileStorage() *FileStorage {
	return defaultFs
}

func (fs *FileStorage) QueryUserByChatid(ctx context.Context, chatid uint64) (*models.UserInfo, bool) {
	fs.lck.RLock()
	defer fs.lck.RUnlock()
	uinfo, ok := fs.indexByChatid[chatid]
	if !ok {
		return nil, false
	}
	return uinfo, true
}

func (fs *FileStorage) DeleteByChatid(ctx context.Context, chatid uint64) error {
	fs.lck.Lock()
	defer fs.lck.Unlock()
	if _, ok := fs.indexByChatid[chatid]; !ok {
		return fmt.Errorf("not found")
	}
	delete(fs.indexByChatid, chatid)
	return fs.save()
}

func (fs *FileStorage) copy(uinfo *models.UserInfo) *models.UserInfo {
	return &models.UserInfo{
		Code:   uinfo.Code,
		Chatid: uinfo.Chatid,
	}
}

func (fs *FileStorage) UpdateUser(ctx context.Context, uinfo *models.UserInfo) error {
	fs.lck.Lock()
	defer fs.lck.Unlock()
	uinfo = fs.copy(uinfo)
	if _, ok := fs.indexByChatid[uinfo.Chatid]; !ok {
		return fmt.Errorf("not found")
	}
	fs.indexByChatid[uinfo.Chatid] = uinfo
	return fs.save()
}

func (fs *FileStorage) NewUser(ctx context.Context, chatid uint64, code string) error {
	fs.lck.Lock()
	defer fs.lck.Unlock()
	uinfo := &models.UserInfo{
		Chatid: chatid,
		Code:   code,
	}
	if _, ok := fs.indexByChatid[uinfo.Chatid]; ok {
		return fmt.Errorf("already exist")
	}
	fs.indexByChatid[uinfo.Chatid] = uinfo
	return fs.save()
}

func (fs *FileStorage) save() error {
	storage := &models.FileStorage{}
	for _, item := range fs.indexByChatid {
		storage.Users = append(storage.Users, &models.UserInfo{
			Chatid: item.Chatid,
			Code:   item.Code,
		})
	}
	data, err := json.Marshal(storage)
	if err != nil {
		return err
	}
	tmp := fs.file + ".tmp"
	if err := ioutil.WriteFile(tmp, data, 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(fs.file, data, 0644); err != nil {
		return err
	}
	return nil
}

func (fs *FileStorage) initFs() error {
	data, err := ioutil.ReadFile(fs.file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	storage := &models.FileStorage{}
	if err := json.Unmarshal(data, storage); err != nil {
		return err
	}
	for _, item := range storage.Users {
		fs.indexByChatid[item.Chatid] = item
	}
	return nil
}
