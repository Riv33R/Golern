package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var rstring string

func main() {

	var id_number = 0
	//Псевдорандом
	text := make(map[int]string)
	text[1] = "jhkhkjhkjhkhk"
	text[2] = "cxvbxcvbxcvbxcvbcvm"
	text[3] = "nlkbfdlnkdfl"
	text[4] = "kldlewlknm;,ewr"
	text[5] = "jdfjorej;;jk"
	text[6] = "g63454fg544fg"
	text[7] = "df798d54f6dfg4312fg"
	text[8] = "d3f2g684654687"
	text[9] = "rtgfhyhj54654ytj"
	text[10] = "ghjtghj4g65h4j5g4"
	text[11] = "ghj645ty4j687f45b"
	text[12] = "fgjh8543jm6456cv4b"
	text[13] = "cvbdfhtj;lef"
	text[14] = "[poujhbvsdgsgdnfg]"
	text[15] = "bvlkjklejbfj;kl"
	text[16] = "khnsdglkhbv;lkn,s"
	text[17] = "dfg;ljmvc.,m.,ndg"
	text[18] = "cvb;lmj.,mdfgj.,mvc"
	text[19] = "cvbcvbkjn.;sdf"
	text[20] = "pouwerklhkgf;lk"
	text[21] = "ewoilkjblvkjn;nmr"
	text[22] = "vbnpojkl;edrgkj"
	text[23] = "dfklhlvblkjdegf"
	text[24] = "dpoj;lknjgflhknvc"
	text[25] = "ttuiykketdye6rf"
	text[26] = "thjgfnmkvbklj"
	text[27] = "mnk;;lmk;lvkbjn;lm"
	text[28] = "vbn;lkf;lmj;vlbn"
	text[29] = "v[oiedrjlgpdfjg;ld"
	text[30] = "sdgftjghlghmftyjerfgn"


	for {

		//Веб интерфейс
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			//Рандомайзер id
			rand.Seed(time.Now().UnixNano())
			id_number = rand.Intn(30)


			rstring = text[id_number]
			id_number := int(id_number)


			//json объект
			text_map := map[string]interface{}{"id": id_number, "feeling": rstring}
			web_print, err := json.MarshalIndent(text_map, "", "  ")
			if err != nil {
				fmt.Println("error:", err)}
			//fmt.Println(w, string(web_print))//Вывод в уонсоль

			//Вывод на веб страницу
			fmt.Fprint(w, string(web_print))


		})
		http.ListenAndServe(":80", nil)

		}

}