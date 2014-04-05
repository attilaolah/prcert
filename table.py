"""Draw a table of smallest prime factors."""
# coding: utf-8
import sys


def print_table():
    """Print a table based on data from the given file."""
    lines = sys.stdin.readlines()
    if len(sys.argv) > 1:
        rows = int(sys.argv[1])
    else:
        rows = len(lines)
    table = [None for i in range(rows)]
    for line in lines:
        _, prime, mod = [int(i) for i in line.strip().split()]
        for i in range(prime-(mod or prime), len(table), prime):
            if table[i] is None:
                table[i] = prime
    width_1 = len(str(rows))
    width_2 = len(str(max(table)))
    template_0 = ' 10⁹⁹⁹⁹⁹⁹⁹⁹   {} | {{:{}d}}'.format(
        ' '*width_1, width_2)
    template_m = ' 10⁹⁹⁹⁹⁹⁹⁹⁹ + {{:{}d}} |'.format(width_1)
    template_n = template_m + ' {{:{}d}}'.format(width_2)
    print ' n ' + ' '*width_1 + '            | ' + 'p'
    print '-' * (width_1+13+2) + '+' + '-' * (width_2+2)
    for i, row in enumerate(table):
        if row:
            if i == 0:
                print template_0.format(row)
            else:
                print template_n.format(i, row)
        else:
            print template_m.format(i)


if __name__ == '__main__':
    print_table()
