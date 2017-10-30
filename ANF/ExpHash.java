/**
 * Compute exponentially distributed hash values (in range 0..32)
 * @author Rasmus Pagh
 * @version 2016.10.18
 */

import java.util.*;

public class ExpHash {

	private static int[] a = {0x5e19b580,0x2b8f2f46,0x01dbee88,0x178439ae,0xe8b8434d,0x45fa4636,0xbb9c5c8c,0x5bdd6e67,0x930ae839,0x36528b7f,0x4fca205e,0xc50b7622,0xb0a63706,0x93ed56cf,0x81852045,0x21b5e7b9,0x16bbb5a7,0xb4837ca5,0xbd49dc89,0xf75c59ca,0x8e11761b,0xf84d7199,0x39e3ef49,0x9e4a936a,0x8d3842f1,0xf499e83d,0xd1431416,0xaca6d6ca,0x38582a69,0xe32a7dc3,0x2faf957e,0xa766a21e};

	public static int hash(int x) {
		for (int value = 0; value<a.length; value++) {
			int dot = Long.bitCount(a[value] & x) & 1;
			if (dot!=0) return value+1;
		}
		return a.length;
	}

}
