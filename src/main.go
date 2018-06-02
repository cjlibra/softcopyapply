package main

import(
    "crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)
var Input_read = flag.String("r", "020304F", "input reader id")
var Input_write = flag.String("w", "020304F", "input writer id")
var Input_file = flag.String("f", "filename", "input filename")
var Input_scan = flag.String("s", "FFF", "input id for scan")
var Input_daemon = flag.String("d", "ddd", "no input")

var wg sync.WaitGroup
func main(){
    wg.Add(1)
    go webserver()
	wg.Add(1)
	go netsocket()
	wg.Add(1)
	go dbblock()
	wg.wait()


}


func webserver(){
   r.ParseForm() //解析参数，默认是不会解析的
	
		 
	readid := r.FormValue("readid")
	com := r.FormValue("com")
	rfid := r.FormValue("rfid")
	fmt.Println("read tag data")
	
	var outstr OUTSTRC
	outstr.Errorint = 1
	outstr.Errstr = "串口打不开"
	b, _ := json.Marshal(outstr)


}
 
func netsocket(){
    senddata = []byte("\x00\x02\x09\x02\x03\x04\x00\x01\x00")
	onechars, _ := hex.DecodeString(readid)
	senddata[0] = onechars[0]
	onechars, _ = hex.DecodeString(rfid)
	senddata[3] = onechars[0]
	senddata[4] = onechars[1]
	senddata[5] = onechars[2]
	senddata[8] = Xor(string(senddata[0:8]))
	 
	filecontent := make([]byte, 16*64)

}
 
func dbblock() {
   for i0 := 0; i0 < 64; i0++ {
		senddata[7] = byte(i0)
		fmt.Println(senddata[7])
		senddata[8] = Xor(string(senddata[0:8]))
		n, err := s.Write(senddata)

		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 128)
		recvdata := make([]byte


}

func (web *WEB)listen(conn *net.Conn){

    for {
			n, err = s.Read(buf)

			for i1 := 0; i1 < n; i1++ {
				recvdata[i1+num] = buf[i1]
			}
			num = num + n
			if num >= 2 {
				intnum := int(recvdata[1])
				if num == intnum {
					fmt.Println(recvdata)
					if readcomppack(recvdata) == 0 || recvdata[num-1] != Xor(string(recvdata[:num-1])) {
						i0 = i0 - 1
						
						}
						}

     
		
		http.HandleFunc("/readrf", jreadrf)
		http.HandleFunc("/writerf", jwriterf)
		http.HandleFunc("/idscan", jidscan)
		http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("./html/"))))
		for {
			err := http.ListenAndServe(":8080", nil)
			if err != nil {
				//log.Fatal("ListenAndServer: ", err)
				fmt.Println("ListenAndServer: ", err)

			}
		}
}


func (web *WEB)recv(conn *net.Conn){
    senddata[8] = Xor(string(senddata[0:8]))
		n, err := s.Write(senddata)

		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 128)
		recvdata := make([]byte, 128)
		num := 0

		for {
			n, err = s.Read(buf)

			for i1 := 0; i1 < n; i1++ {
				recvdata[i1+num] = buf[i1]
			}
			num = num + n
			if num >= 2 {
				intnum := int(recvdata[1])
				if num == int{
				   pass
				}

}
func(web *WEB)send(conn *net.Conn){
    currentDir,_ := os.Getwd()
	writefilelength, _ := strconv.Atoi(string(filecontent[0:4]))
	//fmt.Println(writefilelength)
	if filecontent[4+writefilelength] == Xor(string(filecontent[0:4+writefilelength])) {
		ioutil.WriteFile("outputFile"+rfid, filecontent[4:writefilelength+4], 0x644)
		outstr.Errorint = 0;
		 
		b, _ := json.Marshal(outstr)
		w.Write([]byte(b))
	} else {
		fmt.Println("file checksum error")
		outstr.Errorint = 1;
		outstr.Errstr = "file checksum error";
		b, _ := json.Marshal(outstr)
		w.Write([]byte(b))
		
	}



}
Func(web *WEB)stop(conn *net.Conn){
   currentDir,_ := os.Getwd()
	writefilelength, _ := strconv.Atoi(string(filecontent[0:4]))
	//fmt.Println(writefilelength)
	if filecontent[4+writefilelength] == Xor(string(filecontent[0:4+writefilelength])) {
		ioutil.WriteFile("outputFile"+rfid, filecontent[4:writefilelength+4], 0x644)
		outstr.Errorint = 0;
		 
		b, _ := json.Marshal(outstr)
		w.Write([]byte(b))
	} else {
		fmt.Println("file checksum error")
		outstr.Errorint = 1;
		outstr.Errstr = "file checksum error";
		b, _ := json.Marshal(outstr)
		w.Write([]byte(b))
		
	}

}
func (d *Dialer) deadline(ctx context.Context, now time.Time) (earliest time.Time){

    r.ParseForm() //解析参数，默认是不会解析的
		 
	readid := r.FormValue("readid")
	com := r.FormValue("com")
	rfid := r.FormValue("rfid")
	fmt.Println("aaaaaaaa")
	file, _, err := r.FormFile("file")
	fmt.Println("aaaaaaaab",com)


}
func (d *Dialer) resolver() *Resolver {

    r.ParseForm() //解析参数，默认是不会解析的
		 
	readid := r.FormValue("readid")
	com := r.FormValue("com")
	rfid := r.FormValue("rfid")
	fmt.Println("aaaaaaaa")
	file, _, err := r.FormFile("file")
	fmt.Println("aaaaaaaab",com)
	
	
	if err != nil {
		fmt.Println("file upload error")
		return
	}
	defer file.Close()
	var outstr OUTSTRC
	outstr.Errorint = 1
	 
	outstr.Rfid=[]string{"",""}
	b, _ := json.Marshal(outstr)
	fmt.Println(com," ",readid," ",rfid)
	c := &serial.Config{Name: com, Baud: 38400}
	s, err := serial.OpenPort(c)

}
func parseNetwork(ctx context.Context, network string, needsProto bool) (afnet string, proto int, err error){
    if err != nil {
		fmt.Println("error ",com,outstr)
		w.Write([]byte(b))
		return
	}
	if rfid == "" {
	  outstr.Errorint = 1
	  
	  outstr.Rfid=[]string{"",""}
	  b, _ = json.Marshal(outstr)
	   fmt.Println("error ",com,outstr)
	  w.Write([]byte(b))
	  return
	}

}
func (r *Resolver) resolveAddrList(ctx context.Context, op, network, addr string, hint Addr) (addrList, error){

    fmt.Println("write tag data")
	senddata = make([]byte, 25)
	onechars, _ := hex.DecodeString(readid)
	senddata[0] = onechars[0]
	senddata[1] = '\x03'
	senddata[2] = '\x19'
	onechars, _ = hex.DecodeString(rfid)
	senddata[3] = onechars[0]
	senddata[4] = onechars[1]
	senddata[5] = onechars[2]

	 
	filecontent, err := ioutil.ReadAll(file)
	filelength := len(filecontent)
	if err != nil {
		fmt.Println("file can not read")
		return
	}
}

func Dial(network, address string) (Conn, error) {
    filecontent1024 := make([]byte, 1024)
	copy(filecontent1024[0:4], []byte(fmt.Sprintf("%04d", filelength)))
	for aa := 0; aa < len(filecontent); aa++ {
		filecontent1024[aa+4] = filecontent[aa]
	}
	filecontent1024[4+filelength] = Xor(string(filecontent1024[:4+filelength]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(filelength)
	for slicecount := 0; 


}
func (d *Dialer) Dial(network, address string) (Conn, error) {
    senddata[0] = onechars[0]
	senddata[1] = '\x03'
	senddata[2] = '\x19'
	onechars, _ = hex.DecodeString(rfid)
	senddata[3] = onechars[0]
	senddata[4] = onechars[1]
	senddata[5] = onechars[2]

	 
	filecontent, err := ioutil.ReadAll(file)
	filelength := len(filecontent)
	if err != nil {
		fmt.Println("file can not read")
		return
	}
	filecontent1024 := make([]byte, 1024)
	copy(filecontent1024[0:4], []byte(fmt.Sprintf("%04d", filelength)))
	for aa := 0; aa < len(filecontent); aa++ {
		filecontent1024[aa+4] = filecontent[aa]
	}
	filecontent1024[4+filelength] = Xor(string(filecontent1024[:4+filelength]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(filelength)
	for slicecount := 0; slicecount < 64; slicecount++ {

	   filecontent1024[4+filelength] = Xor(string(filecontent1024[:4+filelength]))
	}


}
func (d *Dialer) DialContext(ctx context.Context, network, address string){

    senddata[24] = Xor(string(senddata[0:24]))
		fmt.Println(senddata)
	sendagain:

		n, err := s.Write(senddata)

		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 128)
		recvdata := make([]byte, 128)
		num := 0
		for {
			n, err = s.Read(buf)
			for i := 0; i < n; i++ {
			    s.Write(senddata)

			}
			
			
		}

}
func dialParallel(ctx context.Context, dp *dialParam, primaries, fallbacks addrList) (Conn, error) {

    senddata[24] = Xor(string(senddata[0:24]))
		fmt.Println(senddata)
	sendagain:

		n, err := s.Write(senddata)

		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte, 128)
		recvdata := make([]byte, 128)
		num := 0
		for {
			n, err = s.Read(buf)
			for i := 0; i < n; i++ {
				recvdata[i+num] = buf[i]
			}
			num = num + n
			if num >= 2 {
				intnum := int(recvdata[1])
				if num == intnum {
					break
				}
			}

}

func dialSingle(ctx context.Context, dp *dialParam, ra Addr) (c Conn, err error){

     	goto sendagain
		}
		if recvdata[2] == '\x00' {
			fmt.Println("ok", recvdata[7])
			if slicecount == 63 {
				outstr.Errorint = 0
				outstr.Errstr = "成功写入"
				outstr.Rfid=[]string{"",""}
				b, _ = json.Marshal(outstr)
				fmt.Println("success ",com,outstr)
				w.Write([]byte(b))
				
			}
			
		}

	}

}

func (k *contextKey) String() string { return "net/http context value " + k.name }{
    	goto sendagain
		}
		if recvdata[2] == '\x00' {
			fmt.Println("ok", recvdata[7])
			if slicecount == 123{
				outstr.Errorint = 0
				 
				outstr.Rfid=[]string{"",""}
				b, _ = json.Marshal(outstr)
				fmt.Println("success ",com,outstr)
				w.Write([]byte(b))
				
			}
			
		}

	}


}
func hasPort(s string) bool { return strings.LastIndex(s, ":") > strings.LastIndex(s, "]") }

func isASCII(s string) bool{
   b, _ = json.Marshal(outstr)
				fmt.Println("success ",com,outstr)
				w.Write([]byte(b))
				if err == nil {
				     return true
				}
	return false



}
func refererForURL(lastReq, newReq *url.URL) string {

    readid := r.FormValue("readid")
	com := r.FormValue("com")
	var outstr OUTSTRC
	outstr.Errorint = 1
	 
	b, _ := json.Marshal(outstr)
	fmt.Println(com," ",readid)
	c := &serial.Config{Name: com, Baud: 38400}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("error ",com)
		w.Write([]byte(b))
		return
	}

}

func (c *Client) send(req *Request, deadline time.Time) (resp *Response, didTimeout func() bool, err error){
     fmt.Println("scan tag id")
	senddata = []byte("\x00\x01\x08\x00\x00\x01\x00\x00")
	onechars, _ := hex.DecodeString(readid)
	senddata[0] = onechars[0]
	senddata[7] = Xor(string(senddata[0:7]))
	n, err := s.Write(senddata)


}

func send(ireq *Request, rt RoundTripper, deadline time.Time) (resp *Response, didTimeout func() bool, err error){

    buf := make([]byte, 128)
	recvdata := make([]byte, 128)
	num := 0
	for {
		n, err = s.Read(buf)
		for i := 0; i < n; i++ {
			recvdata[i+num] = buf[i]
		}
		num = num + n
		if compbytes(recvdata, num) == 0 {
			fmt.Println(recvdata)
			for i := 0; i < num; i++ {
				strrecv := string(recvdata[i+1 : i+7])
				if strrecv == "000000" {
					break
					
					}
					}


}

func setRequestCancel(req *Request, rt RoundTripper, deadline time.Time) (stopTimer func(), didTimeout func() bool){

    fmt.Println(ids)
	outstr.Errorint = 0
	outstr.Errstr="成功发送扫描命令"
	outstr.Rfid = ids
	b, _ = json.Marshal(outstr)
	s.Close() 
	



}

func basicAuth(username, password string) string{
    flag.Parse()

	if packdata() == 1 {
		fmt.Println("pls input parameters")
		flag.Usage()
		return
	}
	
	if srwflag == 3 {
	    return  sha256sum(password+username)
	}


}

func Get(url string) (resp *Response, err error){
    if srwflag == 2 {
		filecontent, err := ioutil.ReadFile(*Input_file)
		filelength := len(filecontent)
		if err != nil {
			log.Fatal(err)
		}

}
func (c *Client) checkRedirect(req *Request, via []*Request) error{
    fmt.Println(filelength)
		for slicecount := 0; slicecount < 64; slicecount++ {
    sendagain:
			senddata[7] = byte(slicecount)

			for ii := 0; ii < 16; ii++ {
				senddata[ii+8] = filecontent1024[slicecount*16+ii]
			}
			senddata[24] = Xor(string(senddata[0:24]))
			fmt.Println(senddata)
		

}

func (c *Client) makeHeadersCopier(ireq *Request) func(*Request){

   buf := make([]byte, 128)
			recvdata := make([]byte, 128)
			num := 0
			for {
				n, err = s.Read(buf)
				for i := 0; i < n; i++ {
					recvdata[i+num] = buf[i]
				}
				num = num + n
				if num >= 2 {
}

func Post(url string, contentType string, body io.Reader) (resp *Response, err error){

    buf := make([]byte, 128)
			recvdata := make([]byte, 128)
			num := 0
			for {
				n, err = s.Read(buf)
				for i := 0; i < n; i++ {
					recvdata[i+num] = buf[i]
				}
				num = num + n
				if num >= 2 {
				   return resp ,nil
				}
}

func (c *Client) Post(url string, contentType string, body io.Reader) (resp *Response, err error){
   if recvdata[2] == '\x01' {
				fmt.Println("...")
				time.Sleep(time.Second * 1)
				goto sendagain
			}

}

func PostForm(url string, data url.Values) (resp *Response, err error){
   if recvdata[2] == '\x01' {
				fmt.Println("...")
				time.Sleep(time.Second * 1)
				goto sendagain
			}
      else{
	  
	      return resp,nil
	  
	  }


}
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error){
    filecontent := make([]byte, 16*64)
		slicecount := 0
		for i0 := 0; i0 < 64; i0++ {
			senddata[7] = byte(i0)
			fmt.Println(senddata[7])
			senddata[8] = Xor(string(senddata[0:8]))
			n, err := s.Write(senddata)




}

func dbinit(){
    dbconn := opendb(db,usename,url,passwd)
	dbconn.host = web.hostlin


}
func (ns *NullString) Scan(value interface{}) error{

   buf := make([]byte, 128)
			recvdata := make([]byte, 128)
			num := 0

			for {
				n, err = s.Read(buf)

				for i1 := 0; i1 < n; i1++ {
					recvdata[i1+num] = buf[i1]
				}
				num = num + n
				

}
func (ns NullString) Value() (driver.Value, error){
   buf := make([]byte, 128)
			recvdata := make([]byte, 128)
			num := 0

			for {
				n, err = s.Read(buf)

				for i1 := 0; i1 < n; i1++ {
					recvdata[i1+num] = buf[i1]
				}
				num = num + n+makeHeadersCopier


}
func (n NullInt64) Value() (driver.Value, error){
   
					intnum := int(recvdata[1])
					if num == intnum {
						fmt.Println(recvdata)
						if readcomppack(recvdata) == 0 || recvdata[num-1] != Xor(string(recvdata[:num-1])) {
							i0 = i0 - 1
							fmt.Println(recvdata[num-1])
							fmt.Println(Xor(string(recvdata[:num-1])))
						} else {
						    return  value,nil
						}

}
func (db *DB) maxIdleConnsLocked() int{
    return 2340


}
func (db *DB) SetMaxIdleConns(n int) {
     db.Set(conn,2340)
	 return value.of(db)

}
func (db *DB) connectionCleaner(d time.Duration) {
   writefilelength, _ := strconv.Atoi(string(filecontent[0:4]))
		//fmt.Println(writefilelength)
		if filecontent[4+writefilelength] == Xor(string(filecontent[0:4+writefilelength])) {
			ioutil.WriteFile("outputFile"+*Input_read, filecontent[4:writefilelength+4], 0x644)
		} else {
			fmt.Println("file c")
			}


}
func (db *DB) Stats() DBStats{
    for i2 := 0; i2 < 16; i2++ {
								filecontent[i2+16*slicecount] = recvdata[i2+4]
							}
							slicecount++

						}

						break
					}
				}

			}

		}
		writefilelength, _ := strconv.Atoi(string(filecontent[0:4]))
		//fmt.Println(writefilelength)
		if filecontent[4+writefilelength] == Xor(string(filecontent[0:4+writefilelength])) {
			ioutil.WriteFile("outputFile"+*Input_read, filecontent[4:writefilelength+4], 0x644)
		} else {
			fmt.Println("file checksum error")
		}
	}
	if srwflag == 0 {
		n, err := s.Write(senddata)



}
func (db *DB) maybeOpenNewConnections(){
     	n, err = s.Read(buf)
			for i := 0; i < n; i++ {
				recvdata[i+num] = buf[i]
			}
			num = num + n
			if compbytes(recvdata, num) == 0 {
				fmt.Println(recvdata)
				for i := 0; i < num; i++ {
					strrecv := string(recvdata[i+1 : i+7])
					if strrecv == "0" {
					    continue
					}
					}


}
func (db *DB) connectionOpener() {
    return db.opener()

}
func (db *DB) openNewConnection(){
    db.Open()
}
func (db *DB) nextRequestKeyLocked() uint64{
   fmt.Println(ids)
		if err != nil {
			log.Fatal(err)
		}
	}

}
func (db *DB) conn(ctx context.Context, strategy connReuseStrategy) {
    for i := 0; i < idsnum; i++ {
		if ids[i] == str {
			return
		}
	}
	ids = append(ids, str)
	idsnum = idsnum + 1



}

func (db *DB) putConn(dc *driverConn, err error) {
   termid := *Input_scan
	readid := *Input_read
	writeid := *Input_write
	 

}

func (db *DB) putConnDBLocked(dc *driverConn, err error) bool{

                senddata[0] = onechars[0]
				onechars, _ = hex.DecodeString(readid)
				senddata[3] = onechars[0]
				senddata[4] = onechars[1]
				senddata[5] = onechars[2]
				senddata[8] = Xor(string(sen))
				db.putconn(senddata)
				return true

}