import io = std.stdio;
import std.path;

void main() {
   auto s = buildNormalizedPath("south", "north");
   io.writeln(s == `south\north`);
}
