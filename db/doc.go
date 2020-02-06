package db

import(
    "encoding/json"
)


type Doc map[string]interface{}


func NewDoc(data []byte)(*Doc,error){
    d := new(Doc)
    err := json.Unmarshal(data,d)
    return d, err
}

func (d *Doc)GetIn(path []string)(ret []interface{}){
    var thing interface{} = d
	// Get into each path segment
	for i, seg := range path {
		if aMap, ok := thing.(map[string]interface{}); ok {
			thing = aMap[seg]
		} else if anArray, ok := thing.([]interface{}); ok {
			for _, element := range anArray {
				ret = append(ret, GetIn(element, path[i:])...)
			}
			return ret
		} else {
			return nil
		}
	}
	switch thing := thing.(type) {
	case []interface{}:
		return append(ret, thing...)
	default:
		return append(ret, thing)
	}
}
