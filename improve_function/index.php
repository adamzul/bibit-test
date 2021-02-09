<?php
function findFirstStringInBracket($str)
{
    if (strlen($str) > 0) {
        $open = strpos($str, '(');
        $close = strpos($str, ')');
       
        if ($open != "" && $close != "" && $open < $close) {
            $length = $close - $open;
            return substr($str, $open+1, $length-1);
        } 
    } 
    return '';
}


echo findFirstStringInBracket("adam(bibit)zulkarnain");