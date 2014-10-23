#! /bin/bash
#set -e

src="gsl-1.16" #original source
dst="cleaned"

rm -rf $dst
mkdir $dst


# in main src dir
# replace '#include <xxx.h>' by '#include "xxx.h"'
# and '#include <gsl/xxx.h>' by '#include "gsl_xxx.h"'
d=$src
for f in $d/*.c $d/*.h; do
		sed -r 's/#include <([a-z|\/|_]*\.h)>/#include "\1"/g' $f | sed 's/#include "gsl\//#include "gsl_/g' > $dst/$(basename $f)
done

# now for src subdirs:
dirs="gsl blas"

for d in $dirs; do
	d=$src/$d
	for f in $d/*.c $d/*.h; do
			sed -r 's/#include <([a-z|_]*)\/([a-z|_]*)\.h>/#include "\1_\2\.h"/g' $f | sed -r 's/#include <([a-z|\/|_]*\.h)>/#include "\1"/g' > $dst/$(basename $d)_$(basename $f)
	done
done



#(cd $dst && gcc -c *.c)
