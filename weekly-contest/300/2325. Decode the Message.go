package contest

func decodeMessage(key string, message string) string {
	m := make(map[byte]byte)
	m[32] = 32
	var count byte = 97
	for i := 0; i < len(key); i++ {
		if _, ok := m[key[i]]; !ok {
			m[key[i]] = count
			count++
		}
	}
	res := ""
	for i := 0; i < len(message); i++ {
		res += string(m[message[i]])
	}
	return res
}
