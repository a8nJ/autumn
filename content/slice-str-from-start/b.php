<?php
extension_loaded('mbstring') or die('mbstring');
$s1 = '♠♣♥♦';
$s2 = mb_substr($s1, 0, 1);
var_dump($s2 == '♠');
