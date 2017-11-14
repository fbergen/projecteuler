p98:
	cat p098_words.txt | sed 's/"//g' | sed 's/,/\n/g' > p098_words_lines.txt
