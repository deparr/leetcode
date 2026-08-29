class Solution:
    def discountPrices(self, sentence: str, discount: int) -> str:
        words: list[str] = sentence.split(" ")
        percent: float = (100 - discount) / 100
        for i, w in enumerate(words):
            if w[0] == "$" and w[1:].isdigit():
                words[i] = f"${float(w[1:]) * percent:0.2f}"

        return " ".join(words)

