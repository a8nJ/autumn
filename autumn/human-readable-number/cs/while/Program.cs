using System;

class Program {
   static string HumanReadable(double len) {
      int order = 0;
      while (len >= 1024) {
         order++;
         len /= 1024;
      }
      return String.Format("{0:0.###}", len) + ' ' + " KMG"[order];
   }

   static void Main() {
      var s = HumanReadable(1264);
      Console.WriteLine(s == "1.234 K");
   }
}
