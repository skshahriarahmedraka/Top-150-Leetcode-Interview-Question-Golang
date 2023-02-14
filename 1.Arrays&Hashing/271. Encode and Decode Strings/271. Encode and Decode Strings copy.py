"""
@param: strs: a list of strings
@return: encodes a list of strings to a single string.
"""
def encode(self, strs):
	res = ''
	for s in strs:
		encoded = str(len(s)) + '/' + s
		res += encoded
	return res    """
@param: str: A string
@return: dcodes a single string to a list of strings
"""
def decode(self, str):
	res, i = [], 0
	while i < len(str):
		# For example, 12/abc
		e = i
		while e < len(str) and str[e] != '/':
			e += 1
		size = int(str[i:e])
		word = str[e + 1, e + 1 + size]
		i = e + 1 + size
		res.append(word)
	return res