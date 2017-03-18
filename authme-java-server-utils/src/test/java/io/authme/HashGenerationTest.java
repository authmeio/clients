package io.authme;

import java.util.HashMap;
import java.util.Map;

import static org.junit.Assert.assertEquals;

public class HashGenerationTest {
  @org.junit.Test
  public void generateHash() throws Exception {


    Map<String, Object> requestMap = new HashMap<String, Object>();

    requestMap.put("test", "hello");

    String apiKey = "test";
    String apiSecret = "secret";


    String generatedHash = HashUtils.generateHash(requestMap, apiKey, apiSecret);
    assertEquals(generatedHash, "292115573432d504b797a836e3a1936ea5ce9ef61bd90c8096dbd86c449d3d75");
  }

}