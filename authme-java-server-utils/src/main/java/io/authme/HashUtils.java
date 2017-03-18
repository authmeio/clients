package io.authme;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;


class HashUtils {

  private static final String SEPARATOR = "|";
  private static final String HASH_KEY = "hash";
  private static final String SHA_256 = "SHA-256";

  public static void main(String[] args) throws java.lang.Exception {
    Map<String, Object> re = new HashMap<String, Object>();
    re.put("test", "hello");
    System.out.println(generateHash(re, "test", "secret"));
  }

  public static String generateHash(Map<String, Object> requestMap, String apiKey, String apiSecret) throws NoSuchAlgorithmException {
    String[] keys = requestMap.keySet().toArray(new String[]{});
    Arrays.sort(keys);
    StringBuilder hashString = new StringBuilder(apiKey);

    for (String key : keys) {
      if (HASH_KEY.equalsIgnoreCase(key)) {
        continue;
      }
      Object value = requestMap.get(key);
      hashString.append(SEPARATOR).append(String.valueOf(value));
    }

    hashString.append(SEPARATOR).append(apiSecret);

    MessageDigest digest = MessageDigest.getInstance(SHA_256);
    byte[] hash = digest.digest(hashString.toString().getBytes());
    return bytesToHex(hash);


  }

  private static String bytesToHex(byte[] bytes) {
    StringBuilder result = new StringBuilder();
    for (byte b : bytes) result.append(Integer.toString((b & 0xff) + 0x100, 16).substring(1));
    return result.toString();
  }

}