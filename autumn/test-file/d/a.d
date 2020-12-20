import file = std.file;
import io = std.stdio;

bool testFile(string s) {
   if (! file.exists(s)) {
      return false;
   }
   auto o = file.getAttributes(s);
   return file.attrIsFile(o);
}

void main() {
   auto b = testFile("a.d");
   io.writeln(b);
}
