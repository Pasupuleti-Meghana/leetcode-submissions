class Solution {
    public boolean isAnagram(String s, String t) {

        // Step 1: If lengths are different, they can't be anagrams
        if (s.length() != t.length()) {
            return false;
        }

        HashMap<Character, Integer> map = new HashMap<>();

        // Step 2: Count characters from string s
        for (char ch : s.toCharArray()) {
            map.put(ch, map.getOrDefault(ch, 0) + 1);
        }

        // Step 3: Decrease count using string t
        for (char ch : t.toCharArray()) {
            if (!map.containsKey(ch)) {
                return false;
            }
            map.put(ch, map.get(ch) - 1);

            if (map.get(ch) == 0) {
                map.remove(ch);
            }
        }

        // Step 4: If map is empty, it's an anagram
        return map.isEmpty();
    }
}