<?php
$a_src = ['May', 'June'];
$f_red = fn ($s_acc, $s_cur) => $s_acc . $s_cur;
$s_acc = '';

foreach ($a_src as $s_cur) {
   $s_acc = $f_red($s_acc, $s_cur);
}

var_dump($s_acc);
