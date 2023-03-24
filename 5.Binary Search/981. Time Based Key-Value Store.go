package main

func main() {
	
}

type TimeMap struct {
    Stamp map[[2]interface{}]string
}


func Constructor() TimeMap {
    return TimeMap{Stamp:make(map[[2]interface{}]string)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.Stamp[[2]interface{}{key,timestamp}] = value
}


func (this *TimeMap) Get(key string, timestamp int) string {
	for i:=timestamp; i>=0; i--{
		if val,ok:=this.Stamp[[2]interface{}{key,i}]; ok{
			return val
		}
	}
	return ""
    
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */