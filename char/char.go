package char

// See: https://scifi.stackexchange.com/questions/137575/is-there-a-list-of-the-symbols-shown-in-the-matrixthe-symbols-rain-how-many
const (
	Unicode = `日ﾊﾐﾋｰｳｼﾅﾓﾆｻﾜﾂｵﾘｱﾎﾃﾏｹﾒｴｶｷﾑﾕﾗｾﾈｽﾀﾇﾍ012345789Z:・."=*+-<>¦｜╌ﾘ       `
	ASCII   = `abcdefghijklmnopqrstuvwxyz0123456789:.";'#=*+-<>|_%^&()!\       `
)

// Get returns the rune at the specified index from either the unicode or ASCII sets.
func Get(i int, useASCII bool) rune {
	if useASCII {
		return []rune(ASCII)[i]
	}
	return []rune(Unicode)[i]
}
