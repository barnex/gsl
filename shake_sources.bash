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
dirs="gsl blas block bspline" 

for d in $dirs; do
	d=$src/$d
	for f in $d/*.c $d/*.h; do
			sed -r 's/#include "([a-z|_]*)\.h"/#include "'$(basename $d)'_\1\.h"/g' $f | \
			sed -r 's/#include "[a-z]*_templates_([o|f|n]*)\.h"/#include "templates_\1\.h"/g'  | \
			sed -r 's/#include "([a-z|_|\.]*)_source.c"/#include "'$(basename $d)'_\1_source.c.inc"/g' | \
			sed -r 's/#include <gsl\/([a-z|_|\.]*)>/#include "gsl_\1"/g'  | \
			sed -r 's/#include <(config|gsl_version|gsl_types)\.h>/#include "\1\.h"/g' > $dst/$(basename $d)_$(basename $f)
	done
done


rm $dst/*_\*.* # stray files

# some people think it's a good idea to include .c files.
# move them aside so they don't get compiled twice
for f in $dst/*_source.c; do
		mv $f $f.inc
done

rm $dst/*_test_*.c $dst/*_test.c


(cd $dst && gcc -c -w *.c)
