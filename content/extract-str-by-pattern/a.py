import re
s1 = 'Wednesday'
s2 = 'e.'
# example 1
o1 = re.search(s2, s1)
s3 = o1.group()
# example 2
a1 = re.findall(s2, s1)
# print
print(s3, a1)
