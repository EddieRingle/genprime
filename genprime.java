public class genprime {
	private static boolean isprime(long x)
	{
		long lim, y;
		if (x < 2)
			return false;
		if (x == 2)
			return true;
		if (x % 2 == 0)
			return false;
		if (x < 9)
			return true;
		if ((x + 1) % 6 != 0)
			if ((x - 1) % 6 != 0)
				return false;
		lim = (long)(Math.sqrt((double)x) + 1.0f);
		for (y = 3; y < lim; y += 2)
		{
			if (x % y == 0)
				return false;
		}
		return true;
	}

	private static long genprime(long max)
	{
		long count = 0,
		     current = 1;
		while (count < max)
		{
			if (isprime(current))
				count++;
			current++;
		}
		return current - 1;
	}

	public static void main(String[] args)
	{
		long start = Long.parseLong(args[0]),
		stop = Long.parseLong(args[1]) + 1,
		last;
		for ( long x = start; x < stop; x += start )
		{
			long starttime = System.nanoTime();
			last = genprime(x);
			long endtime = System.nanoTime();
			double duration = ((double)endtime - (double)starttime) / 1000000000.0;
			System.out.printf ("Found %8d primes in %10.5f seconds (last was %10d)\n",
				x, duration, last);
		}
	}
}
