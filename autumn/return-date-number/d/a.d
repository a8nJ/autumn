import io = std.stdio;
import std.datetime: Clock;

void main() {
   auto o = Clock.currTime;
   auto n = o.fracSecs.total!"msecs";
   io.writeln(n);
}
