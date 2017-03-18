<?php
/**
 * Created by IntelliJ IDEA.
 * User: artpar
 * Date: 18/3/17
 * Time: 3:37 PM
 */

namespace authme;
require "AuthmeClient.php";

class Test extends \PHPUnit_Framework_TestCase
{


    public function testHashGeneration()
    {

        $request = array();
        $request["test"] = "hello";
        $hash = AuthmeClient::generate_hash($request, "test", "secret");
        if ($hash != "292115573432d504b797a836e3a1936ea5ce9ef61bd90c8096dbd86c449d3d75") {
            print "Hash is $hash and not 292115573432d504b797a836e3a1936ea5ce9ef61bd90c8096dbd86c449d3d75";
            throw new \Exception("hash mismatch");
        }
        print "Hash: $hash";
    }
}
