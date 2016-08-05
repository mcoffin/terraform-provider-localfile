package localfile

import (
	"github.com/hashicorp/terraform/helper/schema"
	"os"
  "io/ioutil"
)

func resourceLocalFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceLocalFileCreate,
		Read: resourceLocalFileRead,
		Update: nil,
		Delete: resourceLocalFileDelete,
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Exists: resourceLocalFileExists,
	}
}

type localfileCfg struct {
	Path string
	Content string
}

func newLocalFileCfg(rd *schema.ResourceData) localfileCfg {
	return localfileCfg{
		Path: rd.Get("path").(string),
		Content: rd.Get("content").(string),
	}
}

func resourceLocalFileCreate(rd *schema.ResourceData, _ interface{}) error {
	config := newLocalFileCfg(rd)
	f, err := os.Create(config.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	rd.SetId(config.Path)

	_, err = f.WriteString(config.Content)
	f.Sync()
	return nil
}

func resourceLocalFileDelete(rd *schema.ResourceData, _ interface{}) error {
	config := newLocalFileCfg(rd)
	return os.Remove(config.Path)
}

func resourceLocalFileRead(rd *schema.ResourceData, _ interface{}) error {
	config := newLocalFileCfg(rd)

  x, err := ioutil.ReadFile(config.Path)
  if err != nil {
    return err
  }

  rd.SetId(config.Path)
  rd.Set("content", string(x))

  return nil


}

func resourceLocalFileExists(rd *schema.ResourceData, _ interface{}) (bool, error) {
	config := newLocalFileCfg(rd)

	_, err := os.Stat(config.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}
