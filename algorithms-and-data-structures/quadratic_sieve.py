"""
Se ha decidido utilizar el algoritmo 'Quadratic Sieve', breve descripcion:
The algorithm attempts to set up a congruence of squares modulo n (the integer to be factorized), which often leads to a factorization of n.
The algorithm works in two phases: the data collection phase, where it collects information that may lead to a congruence of squares; 
and the data processing phase, where it puts all the data it has collected into a matrix and solves it to obtain a congruence of squares. 
The data collection phase can be easily parallelized to many processors, but the data processing phase requires large amounts of memory, 
and is difficult to parallelize efficiently over many nodes or if the processing nodes do not each have enough memory to store the whole matrix. 
The block Wiedemann algorithm can be used in the case of a few systems each capable of holding the matrix.
The naive approach to finding a congruence of squares is to pick a random number, square it, divide by n and hope the least non-negative remainder 
is a perfect square. For example, {\displaystyle 80^{2}\equiv 441=21^{2}{\pmod {5959}}}{\displaystyle 80^{2}\equiv 441=21^{2}{\pmod {5959}}}. 
This approach finds a congruence of squares only rarely for large n, but when it does find one, more often than not, the congruence is nontrivial 
and the factorization is complete. This is roughly the basis of Fermat's factorization method.
The quadratic sieve is a modification of Dixon's factorization method.
The general running time required for the quadratic sieve (to factor an integer n) is
{\displaystyle e^{(1+o(1)){\sqrt {\ln n\ln \ln n}}}=L_{n}\left[1/2,1\right]}e^{{(1+o(1)){\sqrt  {\ln n\ln \ln n}}}}=L_{n}\left[1/2,1\right]
in the L-notation.[2]
The constant e is the base of the natural logarithm.
"""
from math import log
from tools import prime_sieve, next_prime, modular_sqrt, legendre_symbol, reduceRowEchelonForm, mod_inv, isqrt
from fractions import gcd


def algebra(factorbase, smooths, settings):
    n = settings
    mvector = [createVector(x[1][0], factorbase) for x in smooths]
    factorbase = [-1]+factorbase
    hvector = [1 << i for i in xrange(len(mvector))]
    reduceRowEchelonForm(mvector, hvector, len(factorbase))
    nulcols = [hvector[x] for x in xrange(len(mvector)) if mvector[x] == 0]
    for nc in nulcols:
        lhs = 1
        rhs = [0]*len(factorbase)
        rhspr = 1
        for i in xrange(0, len(smooths)):
            if nc & (1 << i):
                lh, (rh, ra) = smooths[i]
                lhs *= lh
                rhspr *= ra
                if rh < 0:
                    rhs[0] += 1
                for j in xrange(1, len(factorbase)):
                    while rh % factorbase[j] == 0:
                        rh /= factorbase[j]
                        rhs[j] += 1
        for j in xrange(0, len(factorbase)):
            rhspr *= pow(factorbase[j], rhs[j] >> 1)
        g = gcd(rhspr-lhs, n)
        if g != 1 and g != n:
            return g
    return None


'''
Given smooths, create matrix to find null spaces mod 2 and
find all possible divisors.
'''


def createVector(n, f):
    a = 0
    lg = len(f)-1
    if n < 0:
        a |= 2 << lg
        n = -n
    for i, p in enumerate(f):
        if n % p == 0:
            c = 0
            while n % p == 0:
                n = int(n/p)
                c += 1
            if c & 1:
                a |= 1 << (lg-i)
    return a


"""
Main qs algorithm
Inputs:
    - n: is the numbers from which we want to get the biggest factor
    - verbose: if we activate the flag 'verbose' we will tell the program to print
                through terminal the step by step information provided by the qs algo.
Output:
    This qs algorith will return the first factor found by the algo.
"""


def qs(n, verbose=0):
    if verbose:
        print "Factoring a", int(log(n, 10)+1), "digit number"
    root2n = isqrt(2*n)
    bound = int(5*log(n, 10)**2)

    factorbase = [2] + \
        [x for x in prime_sieve(bound) if legendre_symbol(n, x) == 1]
    if verbose:
        print "Largest Prime Factor used is", factorbase[-1]
    tsqrt = []
    tlog = []
    for p in factorbase:
        ms = int(modular_sqrt(n, p))
        tsqrt.append(ms)
        tlog.append(log(p, 10))
    xmax = len(factorbase)*60*4
    m_val = (xmax * root2n) >> 1
    thresh = log(m_val, 10) * 0.735
    # ignore small primes
    min_prime = int(thresh*3)
    fudge = sum(tlog[i] for i, p in enumerate(factorbase) if p < min_prime)/4
    thresh -= fudge

    roota = int(isqrt(root2n/xmax))
    roota = max(3, roota+int(roota & 1 == 0))
    smooths = []
    partials = {}
    polynomials = 0

    while 1:
        while 1:
            roota = next_prime(roota)
            if legendre_symbol(n, roota) == 1:
                break
        polynomials += 1
        a = int(roota*roota)
        b = modular_sqrt(n, roota)
        b = int((b-(b*b-n)*mod_inv(2*b, roota)) % a)
        c = int((b*b-n)/a)
        sievesize = 1 << 15
        s1 = {}
        s2 = {}
        # set up values
        for i, p in enumerate(factorbase):
            ainv = pow(a, p-2, p)
            sol1 = ainv*(tsqrt[i]-b) % p
            sol2 = ainv*(-tsqrt[i]-b) % p
            sol1 -= ((xmax+sol1)/p)*p
            sol2 -= ((xmax+sol2)/p)*p
            s1[p] = int(sol1+xmax)
            s2[p] = int(sol2+xmax)
        # segmented sieve
        for low in xrange(-xmax, xmax+1, sievesize+1):
            high = min(xmax, low+sievesize)
            size = high - low
            S = [0.0]*(size+1)
            # sieve segment
            for i, p in enumerate(factorbase):
                if p < min_prime:
                    continue
                sol1 = s1[p]
                sol2 = s2[p]
                logp = tlog[i]
                while sol1 <= size or sol2 <= size:
                    if sol1 <= size:
                        S[sol1] += logp
                        sol1 += p
                    if sol2 <= size:
                        S[sol2] += logp
                        sol2 += p
                s1[p] = int(sol1-size-1)
                s2[p] = int(sol2-size-1)
            # check segment for smooths
            for i in xrange(0, size+1):
                if S[i] > thresh:
                    x = i+low
                    tofact = nf = a*x*x+2*b*x+c
                    if nf < 0:
                        nf = -nf
                    for p in factorbase:
                        while nf % p == 0:
                            nf = int(nf/p)
                    if nf == 1:
                        smooths += [(a*x+b, (tofact, roota))]
                    # check for 1 big factor
                    elif nf in partials:
                        pairv, pairvals = partials[nf]
                        smooths += [((a*x+b)*pairv, (tofact *
                                                     pairvals[0], roota*pairvals[1]*nf))]
                        del partials[nf]
                    else:
                        partials[nf] = (a*x+b, (tofact, roota))
        if len(smooths) > len(factorbase):
            break
        if verbose:
            print 100*len(smooths)/len(factorbase), '%', 'using', polynomials, 'polynomials', '\r',
    if verbose:
        print len(smooths), "relations found using", polynomials, "polynomials"
    return algebra(factorbase, smooths, n)


if __name__ == "__main__":
    from time import time
    print('Introduce number: ')
    number = input()
    t1 = time()
    first_prime = qs(number, 1)
    second_prime = number / first_prime
    print "First Prime: {}\nSecond Prime: {}".format(first_prime, second_prime)
    total_time = time()-t1
    print "Time: {}".format(total_time)


"""
RESULTADOS:
	- First:   11111222223333344444555556666677777888889999900073 (50 digitos)
	- Second:  11111222223333344444555556666677777888889999900079 (50 digitos)
"""

# TEST CASES:

# Try: 44798641826506377809 * 60402714084339067253
# Try: 639374344395443 * 730494381013249
# Try: 125687024840911 * 498588031558063
