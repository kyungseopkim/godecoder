package godecoder

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "math"
    "path"
    yaml "gopkg.in/yaml.v2"
    "github.com/kyungseopkim/goarxml"
    "strconv"
)

const mappingTable = "arxml_mapping.yaml"

type VinMap map[string]int32
type MsgMap map[int32]goarxml.Message
type ArxmlMap map[int32]MsgMap
type ArxmlVins struct {
    Name    string   `yaml:"name"`
    Ver     string   `yaml:"version"`
    Vin     []string `yaml:"vin"`
    Default bool     `yaml:"default"`
    Gps     bool     `yaml:"gps"`
}

type ArxmlVinMap struct {
    Arxml []ArxmlVins `yaml:"arxml"`
}

func ToJson(data interface{}) string {
    byteStr, err := json.Marshal(data)
    if err != nil {
        log.Fatalln(err)
    }
    return string(byteStr)
}

func (vin VinMap) String() string {
    return ToJson(vin)
}

func (msgs MsgMap) String() string {
    return ToJson(msgs)
}

func (arxmls ArxmlMap) String() string {
    return ToJson(arxmls)
}

func ToYaml(data interface{}) string {
    byteStr, err := yaml.Marshal(data)
    if err != nil {
        log.Fatalln(err)
    }
    return string(byteStr)
}

func (vins ArxmlVins) String() string {
    return ToYaml(vins)
}

func (vinmap ArxmlVinMap) String() string {
    return ToYaml(vinmap)
}

func (msgMap MsgMap) FromMessages(msgs []goarxml.Message) {
    for _, msg := range msgs {
        msg.SortByStartbit()
        msgMap[msg.Id] = msg
    }
}

func (dbc ArxmlMap) FromResource(arxmlMap ArxmlVinMap, resource string) {
    for _, arxml := range arxmlMap.Arxml {
        msgMap := make(MsgMap)
        msgMap.FromMessages(arxml.GetMsg(resource))
        ver, _ := strconv.ParseInt(arxml.Ver, 10, 32)
        dbc[int32(ver)] = msgMap
    }
}

func (vins ArxmlVins) GetMsg(resource string) []goarxml.Message {
    fileName := path.Join(resource, "arxml", vins.Name)
    return goarxml.Parse(fileName)
}

func readArxmlVinsFromFile(filePath string) ArxmlVinMap {
    tablePath := path.Join(filePath, "arxml", mappingTable)
    bytes, err := ioutil.ReadFile(tablePath)
    if err != nil {
        log.Fatalln(err)
    }
    var yamlContents ArxmlVinMap
    err = yaml.Unmarshal(bytes, &yamlContents)
    if err != nil {
        log.Fatalln(err)
    }
    return yamlContents
}

func (vinmap VinMap) GetFromArxmlVins(vins ArxmlVinMap) {
    for _, arxml := range vins.Arxml {
        for _, vin := range arxml.Vin {
            ver, _ := strconv.ParseInt(arxml.Ver, 10, 32)
            vinmap[vin] = int32(ver)
        }
    }
}

func GetMappingTables(resources string) (VinMap, ArxmlMap, MsgMap) {
    arxml := readArxmlVinsFromFile(resources)
    vinMap := make(VinMap)
    vinMap.GetFromArxmlVins(arxml)
    arxmlMap := make(ArxmlMap)
    arxmlMap.FromResource(arxml, resources)

    var defaultArxmlVer int32 = 0
    var lowest int32 = math.MaxInt32
    for _, arxml := range arxml.Arxml {
        ver, _ := strconv.ParseInt(arxml.Ver, 10, 32)
        if int32(ver) < lowest {
            lowest = int32(ver)
        }
    	if arxml.Default {
    		defaultArxmlVer = int32(ver)
    		break
	    }
    }

    defaultMap, ok := arxmlMap[defaultArxmlVer]
    if !ok {
        defaultMap, _ = arxmlMap[lowest]
    }
    return vinMap, arxmlMap, defaultMap
}

var (
    Vin2Ver      VinMap
    Ver2Arxml    ArxmlMap
    DefaultArxml MsgMap
)

func ArxmlLoad(resource string) {
    Vin2Ver, Ver2Arxml, DefaultArxml = GetMappingTables(resource)
}
