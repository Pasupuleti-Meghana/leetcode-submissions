class Solution {
    public int maxChunksToSorted(int[] arr) {
        int maxsoFar = 0;
        int chunks = 0;
        int n = arr.length;

        for(int i=0; i<n; i++) {
            maxsoFar = Math.max(maxsoFar, arr[i]);

            if (maxsoFar == i) {
                chunks++;
            }
        }

        return chunks;
        
    }
}