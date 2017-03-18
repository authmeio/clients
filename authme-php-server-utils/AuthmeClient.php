<?php
/**
 * Created by IntelliJ IDEA.
 * User: artpar
 * Date: 18/3/17
 * Time: 3:34 PM
 */

namespace authme;


class AuthmeClient
{


    public static function generate_hash($params, $apiKey, $apiSecret)
    {
        $str = self::flatten_array($params);
        $str .= $apiSecret;
        $str = $apiKey . "|" . $str;
        print $str;
        return hash("sha256", $str);
    }


    private static function flatten_array($data)
    {
        if (empty($data)) {
            return "|";
        }
        $str = "";
        ksort($data);
        foreach ($data as $key => $value) {
            if ($key === "Hash") {
                continue;
            } elseif (is_array($value)) {
                $str .= self::flatten_array($value);
//                TODO: remove nulls and empty
            } elseif (is_null($value)) {
                $str .= "null" . "|";
            } elseif (is_bool($value)) {
                $bool_str = $value ? "true" : "false";
                $str .= $bool_str . "|";
            } else {
                $str .= $value . "|";
            }
        }
        return $str;
    }

}