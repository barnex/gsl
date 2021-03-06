#! /bin/bash
#set -e

src="gsl-1.16" #original source
dst="internal"

rm -f $dst/*.c $dst/*.h $dst/*.o

#in main src dir
d=$src
for f in $d/*.c $d/*.h; do
		sed -r 's/#include <(config|gsl_version|gsl_types)\.h>/#include "\1\.h"/g' $f | \
		sed -r 's/#include <gsl\/([a-z|_]*)\.h>/#include "gsl_\1\.h"/g' > $dst/$(basename $f)
done

# now for src subdirs:
dirs="gsl cblas sys linalg err vector matrix blas complex permutation block fft"

for d in $dirs; do
	d=$src/$d
	for f in $d/*.c $d/*.h; do
			bd=$(basename $d)
			sed -r 's/#include "([a-z|_|0-9]*)\.h"/#include "'$bd'_\1\.h"/g' $f | \
			sed -r 's/#include "[a-z]*_templates_([o|f|n]*)\.h"/#include "templates_\1\.h"/g'  | \
			sed -r 's/#include "([a-z|0-9|_|\.]*)\.c"/#include "'$bd'_\1.c.inc"/g' | \
			sed -r 's/#include <gsl\/([a-z|_|\.]*)>/#include "gsl_\1"/g'  | \
			sed -r 's/#include <(config|gsl_version|gsl_types|build)\.h>/#include "\1\.h"/g' > $dst/$bd'_'$(basename $f)
	done
done


rm $dst/*_\*.* # stray files

rm $dst/gsl-histogram.c

# some people think it's a good idea to include .c files.
# move them aside so they don't get compiled twice
for f in $dst/*_source.c; do
		mv $f $f.inc
done

#quirks
(cd $dst

mv cblas_hypot.c cblas_hypot.c.inc

for f in *_inc.c; do
		mv $f $f.inc;
done;

mv eigen_qrstep.c eigen_qrstep.c.inc

for f in linalg_givens linalg_apply_givens linalg_svdstep; do
	mv $f.c $f.c.inc;
done

for f in fft_real_init.c fft_real_main.c fft_real_pass_n.c fft_real_unpack.c fft_hc_init.c fft_hc_radix*.c fft_real_radix*.c fft_hc_main.c fft_hc_pass_n.c fft_hc_unpack.c fft_c_init.c fft_c_main.c fft_c_pass_n.c fft_c_unpack.c fft*_pass*.c fft_c_radix*.c fft_bitreverse.c fft_urand.c fft_factorize.c; do
	mv $f $f.inc;
done

rm *_test_*.c *_test.c *_tests.c

cp build.h combination_build.h
cp build.h complex_build.h
cp build.h vector_build.h
cp build.h matrix_build.h
cp build.h permutation_build.h
cp build.h sys_build.h


rm gsl-historgram.c
rm gsl-randist.c
)


