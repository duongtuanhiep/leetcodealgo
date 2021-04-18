package amazontest;

import java.io.*;
import java.util.*;

public class Solution {
	// METHOD SIGNATURE BEGINS, THIS METHOD IS REQUIRED
	public ALNode deepCopy(ALNode head) {
		ArrayList<ALNode> arr = new ArrayList<ALNode>();
		while (head != null) {
			ALNode newNode = new ALNode(head.value);
			newNode.value = head.value;
			arr.add(newNode);
			head = head.next;
		}

		for (int i = 0; i < arr.size(); i++) {
			if (i < arr.size() - 1) {
				arr.get(i).next = arr.get(i + 1);
			}
		}

		if (arr.size() > 0) {
			return null;
		}
		return arr.get(0);
	}
	// METHOD SIGNATURE ENDS
}

class ALNode{
	public int value; 
	public ALNode next;
	public ALNode abitrary;
}