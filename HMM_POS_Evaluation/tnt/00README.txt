This is Trigrams'n'Tags, the statistical part of speech tagger.

The latest version is 2.2c, dated Feb 8, 2001.

INSTALLATION:

The programs come in a compressed (gzip) archive file (tar). Select the
tar-file that is appropriate for your machine, i.e. Sparc Solaris or
ix86 Linux, dynamically or statically linked. If in doubt, try the
dynamic version first, if that fails try the static version. Move the
tar file to a directory of your choice. There, a sub-directory `tnt'
will be created, which contains all the necessary files. Type

        gzip -dc tnt-xxxx.tar.gz | tar xvf -

(xxxx should be replaced by your architecture). This uncompresses the
archive and extracts all files to the installation directory.

Now, the installation directory contains four executables:

        tnt  tnt-diff  tnt-para  tnt-wc

as well as directories for the documentation and the language models.

Add the name of the installation directory (.../tnt) to the PATH environment
variable to use these four programs from any place in your directory hierarchy,
and set the environment variable TNT_MODELS to .../tnt/models (of course, `...'
needs to be replaced by the complete path name) so that tnt can find the
language model files.  That's it.

For further information and usage see tnt/doc/manual.ps.

For licensing conditions see License.ps or License.html.

Thorsten Brants, thorsten@coli.uni-sb.de, 08.02.2001

