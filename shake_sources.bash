#! /bin/bash
#set -e

src="gsl-1.16" #original source
dst="cleaned"

rm -f $dst/*.c $dst/*.h $dst/*.o

#in main src dir
d=$src
for f in $d/*.c $d/*.h; do
		sed -r 's/#include <(config|gsl_version|gsl_types)\.h>/#include "\1\.h"/g' $f | \
		sed -r 's/#include <gsl\/([a-z|_]*)\.h>/#include "gsl_\1\.h"/g' > $dst/$(basename $f)
done

# now for src subdirs:
dirs="gsl blas block bspline cblas cdf cheb combination complex" 

for d in $dirs; do
	d=$src/$d
	for f in $d/*.c $d/*.h; do
			bd=$(basename $d)
			sed -r 's/#include "([a-z|_|0-9]*)\.h"/#include "'$bd'_\1\.h"/g' $f | \
			sed -r 's/#include "[a-z]*_templates_([o|f|n]*)\.h"/#include "templates_\1\.h"/g'  | \
			sed -r 's/#include "([a-z|_|\.]*)\.c"/#include "'$bd'_\1.c.inc"/g' | \
			sed -r 's/#include <gsl\/([a-z|_|\.]*)>/#include "gsl_\1"/g'  | \
			sed -r 's/#include <(config|gsl_version|gsl_types|build)\.h>/#include "\1\.h"/g' > $dst/$bd'_'$(basename $f)
	done
done


rm $dst/*_\*.* # stray files

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

rm *_test_*.c *_test.c

cp build.h combination_build.h
cp build.h complex_build.h

)


(cd $dst && gcc -pipe -fPIC -c -w *.c)
