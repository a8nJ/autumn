# example 1
class Date:
   year = 2019
o = Date()

# example 2
import types
o2 = types.SimpleNamespace(year = 2019)

# print
print(o.year == 2019, o2.year == 2019)
