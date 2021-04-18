package amazontest;

class GCD {
	// METHOD SIGNATURE BEGINS, THIS METHOD IS REQUIRED
	public static int generalizedGCD(int arr[]) {
		// INSERT YOUR CODE HERE
		if (arr.length < 2) {
			return arr[0];
		}
		
		int lastGcd = arr[0];
		for (int i =0; i < arr.length-1; i++) {
			lastGcd = gcd(lastGcd, arr[i+1]);
		}
		return lastGcd;
	}
	// METHOD SIGNATURE ENDS
	public static int gcd(int a, int b) {
		if (a==0) {
			return b;
		}
		return gcd(b%a,a);
	}
}