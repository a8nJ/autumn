<?php
$a = ['May', 'June'];

while (true) {
   $n = key($a);
   if ($n === null) {
      break;
   }
   $s = current($a);
   echo $n, "\t", $s, "\n";
   next($a);
}
