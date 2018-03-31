package main

import (
	"math"
	"os"
	"fmt"
	
)

func main(){
	args := os.Args;
	if len(args) < 2 {
		fmt.Println("请输入文件名")
	} else {
		file := args[1];
		readLine(file);
	}
}

func readLine(path string) bool {
	fi,err := os.OpenFile(path,os.O_APPEND|os.O_CREATE, 0644);
	if err != nil {
		panic(err);
	}
	defer fi.Close();
	for {
		var v1,v2,v3 float64;
		_,err := fmt.Fscanln(fi,&v1,&v2,&v3);
		if err != nil{
			break
		}

		computeRate(v1,v2,v3)
	}
	return true;
}

func computeRate(high float64, middle float64, low float64) bool {
	n := high + middle + low;
	//加权平均
	m := 0.1*high + 0.35*middle + 0.55*low;
	var rate float64;
	if n == 0 {
		rate = 1;
	} else {
		rate = m / n;
	}
	var normalRate float64 = (rate)*100;
	
	//加权自然对数
	m1 := 0.55*high + 0.35*middle + 0.1*low;
	var logRate float64;
	if n == 0 {
		logRate = 100 ;
	}else{
		logRate = (1-math.Log1p(m1)/math.Log1p(2*n))*100 ;
	}
	
	//反比例对数
	m2 := 0.704*high + 0.282*middle + 0.0141*low; 
	var overRate float64;
	overRate =  1/math.Log(math.E+m2) * 100; 
	
	fmt.Printf("高危：%d,中危：%d,低危：%d, 加权评分：%.1f, 自然对数评分：%.1f, 反比例对数：%.1f\n",int(high),int(middle),int(low),normalRate, logRate, overRate);
	return true;
}