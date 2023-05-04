package googledrive

import (
	"fmt"
	"io"
	"os"
	"strings"

	LogTool "github.com/adimax2953/log-tool"
	"golang.org/x/net/context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type Google struct {
	Service *drive.Service
}

// Init - 初始化
func (g *Google) Init(serviceAccountFilePath string) error {
	var err error
	g.Service, err = drive.NewService(context.Background(), option.WithServiceAccountFile(serviceAccountFilePath))
	if err != nil {
		LogTool.LogError("Google Drive Init", err)
		return err
	}
	return nil
}

// GetFileList - 取得檔案列表
func (g *Google) GetFileList(pageSize int64) ([]*drive.File, error) {

	r, err := g.Service.Files.List().PageSize(pageSize).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		LogTool.LogError("Unable to retrieve files Error : ", err)
		return nil, err
	}

	return r.Files, nil
}

// ReadFile - 讀取檔案
func (g *Google) ReadFile(fileID string) (*drive.File, error) {

	file, err := g.Service.Files.Get(fileID).Do() //u can call Get(fileID).Download() to download file content
	if err != nil {
		LogTool.LogError("Unable to read files Error : ", err)
		return nil, err
	}
	return file, nil
}

// DownloadFile - 下載檔案
func (g *Google) DownloadFile(filePath, fileName, fileID string) error {

	// 取得檔案內容
	resp, err := g.Service.Files.Get(fileID).Download()
	if err != nil {
		return fmt.Errorf("無法取得檔案內容： %v", err)
	}
	defer resp.Body.Close()

	// 建立本地檔案
	out, err := os.Create(filePath + fileName)
	if err != nil {
		return fmt.Errorf("無法建立本地檔案： %v", err)
	}
	defer out.Close()

	// 複製檔案內容到本地檔案
	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("無法複製檔案內容： %v", err)
	}
	LogTool.LogInfo("已下載檔案：", fileName)

	return nil
}

// CreateDir -建立資料夾
func (g *Google) CreateDir(name string, parents ...string) (*drive.File, error) {

	d := &drive.File{
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  parents,
	}

	file, err := g.Service.Files.Create(d).Do()

	if err != nil {
		LogTool.LogError("Could not create dir: ", err.Error())
		return nil, err
	}

	return file, nil
}

// CreateFile - 建立檔案
func (g *Google) CreateFile(name, mimeType, content, parentID string) (*drive.File, error) {
	contentReader := strings.NewReader(content)

	parents := []string{}
	if parentID != "" {
		parents = []string{parentID}
	}
	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  parents,
	}
	file, err := g.Service.Files.Create(f).Media(contentReader).Do()

	if err != nil {
		LogTool.LogError("Could not create file: ", err.Error())
		return nil, err
	}

	return file, nil
}
