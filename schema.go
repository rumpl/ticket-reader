package main

import "encoding/json"

func SchemaToMap(params any) map[string]any {
	m := map[string]any{}
	buf, err := json.Marshal(params)
	if err != nil {
		return nil
	}

	if err := json.Unmarshal(buf, &m); err != nil {
		return nil
	}

	return m
}
