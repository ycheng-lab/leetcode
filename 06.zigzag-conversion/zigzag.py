
from py_compile import main
from socket import if_nameindex


def zigzag_conversion(text: str, row_number: int) -> str:
    if len(text) <= 1 or len(text) <= row_number:
        return text

    res = []
    pace = 2 * row_number - 2

    # first row
    for i in range(0, len(text), pace):
        res.append(text[i])

    # 2nd ... last - 1 row
    for row in range(1, row_number - 1):
        res.append(text[row])
        for k in range(pace, len(text), pace):        
            res.append(text[k - row])
            if k + row < len(text):
                res.append(text[k + row])

    # last row
    for i in range(row_number - 1, len(text), pace):
        res.append(text[i])

    return "".join(res)


def solve_problem(text: str, row_number: int) -> None:
    res = zigzag_conversion(text, row_number)
    print(f"string: {text}, row: {row_number}, result: {res}")
    if res == "PAHNAPLSIIGYIR":
        print("zigzag conversion was correct")
    else:
        print("Failed to convert")

if __name__ == '__main__':
    solve_problem("PAYPALISHIRING", 3)
