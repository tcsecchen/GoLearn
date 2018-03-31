package main  

import (  
	"time"
  	"net"  
  	"fmt"  
	"os"
	"MyGo/code/network/dns/file"
	"runtime/pprof"
	"log"
)  


func main() {
	f, err := os.OpenFile("./cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	  
	start := time.Now()
	fd := file.Read("subnames_full.txt");
	for _, url := range os.Args[1:]{
		for i := 0;i <= 9; i++ {
			go CheckDNS(url, fd, 2000*i, 2000*i+2000,2000*i)
		} 
		
		time.Sleep(120*time.Second);
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds());
	pprof.StopCPUProfile()		
	f.Close()
}

func CheckDNS(url string, data []string, start int, end int, i int)string{
	timeStart := time.Now()
	for _, prefix := range data[start:end]{			
		prefix = prefix +"."+ url;
		ns, err := net.LookupIP(prefix)
		fmt.Println(i);
		i++;
		fmt.Printf("%.2fs\n",time.Since(timeStart).Seconds());    
		if err != nil {
			 
			//fmt.Fprintf(os.Stdout, "%s\n", err.Error())  
			continue  
		}  
		for _, n := range ns {  
			fmt.Fprintf(os.Stdout, "%s--%s\n", prefix, n);   
		}
		
		fmt.Printf("%.2fs\n",time.Since(timeStart).Seconds());
		
	}
	return "0"
}