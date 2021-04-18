package amazontest;

import java.io.*;
import java.util.*;

class Colony {
	// METHOD SIGNATURE BEGINS, THIS METHOD IS REQUIRED
	public static int[] cellCompete(int[] cells, int days) {
		// INSERT YOUR CODE HERE
		for (; days > 0; days--) {
			int[] state = new int[cells.length];
			for (int i = 0; i < cells.length; i++) {
				if (i == 0) {
					if (cells[i + 1] == 1) {
						state[i] = 1;
					} else {
						state[i] = 0;
					}
				} else if (i == cells.length - 1) {
					if (cells[i - 1] == 1) {
						state[i] = 1;
					} else {
						state[i] = 0;
					}
				} else {
					if (cells[i + 1] != cells[i - 1]) {
						state[i] = 1;
					} else {
						state[i] = 0;
					}
				}
			}
			cells = state;
		}
		return cells;
	}
	// METHOD SIGNATURE ENDS
}