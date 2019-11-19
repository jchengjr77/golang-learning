/* Package stringutil has utility functions for working with strings
 * Made for practice by jchengjr77
 *
 * Simple string reversal code
 */

package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
//  REQUIRES: nothing
//  ENSURES: s' = reversed s
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; j >= len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
