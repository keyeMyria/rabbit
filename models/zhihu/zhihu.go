package zhihu

import (
	"fmt"
	"math/rand"

	"github.com/hunterhug/rabbit/lib"
	"github.com/hunterhug/rabbit/models/blog"
)

var (
	lunzi = "Lunzi"
)

func UpdateConfig() {
	fmt.Println("update config start")
	c := new(blog.Config)
	c.Webinfo = fmt.Sprintf(`
	{
		"1":{"name":"%s","limit":6}
	}
	`, lunzi)
	err := c.Update("Webinfo")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("update config end")
}

func InsertCategory() {
	fmt.Println("insert category start")
	cs := map[int64]string{1: lunzi}
	for k, v := range cs {
		c := new(blog.Category)
		c.Id = k
		c.Title = v + "-T"
		c.Alias = v
		c.Createtime = lib.GetTime()
		c.Status = 1
		c.Image = "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png"
		c.Content = v
		if k == 4 {
			c.Type = 1
		}
		if k > 4 {
			c.Pid = 4
			c.Type = 1
		}
		err := c.Insert()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("insert category end")
}

func InsertPaper() {
	fmt.Println("insert paper start")
	aaa := []string{
		"/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png",
		"/file/image/12/68756e74657268756795ad72d42c7ef1b56c04c66297db1c27.jpeg",
		"/file/image/64/68756e74657268756759fc5ee18fa5210bb76003976900fae9.jpeg",
	}
	k := 140
	for k > 0 {
		k = k - 1
		paper := new(blog.Paper)
		paper.Title = "Test test test test data"
		paper.Status = 1
		paper.Photo = aaa[rand.Intn(3)]
		paper.Descontent = "淭一厘晢臹隒丌蒛霒冘，庳乜砐淈琲葙丌漍厹刌。仈"
		paper.Content = `
So her / Ce are Xi Zan, but what Xiao this Wu may not put up the Yun Zhang ping. These Jie Tong Kang Schrodinger
Jiu Da Chi Yu a Feng Bing Wu Dao tea, Ke Bei Fu Tu cou not drag type go to squall process gas Bureau ze. A Qu% Xiu
Yan Que Yin find not weighted, Bi E Qu Pei cases Guo are not Rou cut. He Ge howl Xi Cang Hu Bo Kuang Nan Xin loop Yi
Fu Ru Chu Le Xie row full of a surname. He got Shan creating Zhi w Gun and Dun Cu Qiong Jian Hu You Yi Wei Ding Hai Li.
After a Xuan Bei Tu yao type seal Zeng ang, Guo Wu You Liu Wen Torr + Quan market dissatisfied. A Xian how prepared and
not Zhu Che Han he Xia, Chu Zhong Chen Yi Yan not falsification of Ji E. A Gong Wu Fei Guo zhe Ju Er through, please.
Guan melancholy that he was Jiang urn donburi. A host Gui Yuncheng Qin TA type. Zhi Ji, Ji type He Xun Zan Mei not Jiao
Dao le. Kun a sincere Sun what taste as a surname Ju Qi, sow not You Si Xiao type Hou tonnes through the. A Yan Zhao
`
		paper.Author = "hunterhug"
		paper.Createtime = lib.GetTime()
		paper.Cid = int64(1)
		if paper.Cid >= 4 {
			paper.Type = 1
		}
		paper.Istop = int64(rand.Intn(2))
		paper.Insert()
	}
}

func InsertRoll() {
	fmt.Println("insert roll start")
	rolls := map[string]string{
		"tuzi":   "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png",
		"me":     "/file/image/37/68756e7465726875673308fd68c821f8fb4180732625ef10ba.png",
		"tuzizi": "/file/image/46/68756e746572687567aadc0c7438bb9e28d2c4eeaa310828e8.png",
		"me1":    "/file/image/37/68756e7465726875673308fd68c821f8fb4180732625ef10ba.png",
	}
	for k, v := range rolls {
		t := new(blog.Roll)
		t.Photo = v
		t.Status = 1
		t.Title = k
		t.Createtime = lib.GetTime()
		t.Insert()
	}
}
