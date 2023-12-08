#!/bin/bash
# .go: ensure virtual environment activation
#      and setup some aliases to be ready to go

export REPO=`pwd`

#===============================================================================
# step 1: setup aliases
#===============================================================================

  alias de="deactivate"
  alias jl="jupyter lab"

  #alias ve="source `pwd`/local/bin/ve.sh"
  #alias ec="bash `pwd`/local/bin/ec"                            # deactivate virtual env.
  #alias po="bash `pwd`/local/bin/po"                            # deactivate virtual env.
  #alias id="python -m idlelib.idle"
  #alias ?="bash `pwd`/local/bin/hlp"

#===============================================================================
# step 2: virtual environment
#===============================================================================

  if [ -d venv ]; then

    source venv/bin/activate

  else
    printf "\x1b[36m"  #  cyan
    echo 'creating virtual environment ...'
    printf "\x1b[0m"   # neutral

    printf "\x1b[33m"  # yellow
    echo 'python3 -m venv venv'
    printf "\x1b[0m"   # neutral

    deactivate 2>/dev/null       # be sure a current venv is deactivated
    python3 -m venv venv
    source venv/bin/activate     # activate virtual environment

    printf "\x1b[33m"  # yellow
    echo 'pip install --upgrade pip'
    printf "\x1b[0m"   # neutral

    pip install --upgrade pip

    printf "\x1b[33m"  # yellow
    echo 'pip install bleak'
    printf "\x1b[0m"   # neutral

    pip install bleak
  fi

#===============================================================================
# step 3: running the sample
#===============================================================================

  printf "\x1b[32m"   # green
  echo 'to run sample => $ python sample/scan.py  # run BLE scanner'
  printf "\x1b[0m"    # neutral
