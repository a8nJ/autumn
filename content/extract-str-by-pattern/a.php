<?php
$s = 'January';
# example 1
preg_match('/a./', $s, $a);
# example 2
preg_match_all('/a./', $s, $a2);
# example 3
preg_match('/a(..)/', $s, $a3);
# example 4
preg_match_all('/a(..)/', $s, $a4);
# print
var_dump($a, $a2, $a3, $a4);
