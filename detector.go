package detector


	

import "strings"
import(
//  log "github.com/sirupsen/logrus"
  )
import "regexp"
import "os/exec"

type NetworkDetector struct {
real_interface_regexp *regexp.Regexp
}

func NewNetworkDetector() NetworkDetector {
  return NetworkDetector{
  	regexp.MustCompile(`^\d+: (eth|wlan)\d+:`),
  }
}

func (d *NetworkDetector) up() (bool, error) {
cmd:=exec.Command("ip","a")
output,err:=cmd.Output()
if err!=nil{return false,err}
lines:=strings.Split(string(output),"\n")
for index,line:=range lines{
if d.real_interface_regexp.MatchString(line){
if strings.Contains(line,"NO-CARRIER"){continue}
if strings.Contains(line,"UP"){return true,nil}
}
index=index
}
return false,nil
}
