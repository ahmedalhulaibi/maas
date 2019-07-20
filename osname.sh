#!/bin/sh
# Detects which OS and if it is Linux then it will detect which Linux
# Distribution.
# derived from https://stackoverflow.com/a/26988390

OS=`uname -s`

if [ "${OS}" = "Linux" ] ; then
    if [ -f /etc/redhat-release ] ; then
        DIST='RedHat'
    elif [ -f /etc/SuSE-release ] ; then
        DIST=`cat /etc/SuSE-release | tr "\n" ' '| sed s/VERSION.*//`
    elif [ -f /etc/mandrake-release ] ; then
        DIST='Mandrake'
    elif [ -f /etc/debian_version ] ; then
        DIST="Debian"

    fi
    if [ -f /etc/UnitedLinux-release ] ; then
        DIST="${DIST}[`cat /etc/UnitedLinux-release | tr "\n" ' ' | sed s/VERSION.*//`]"
    fi
elif
    DIST="unknown"
fi

echo ${DIST}