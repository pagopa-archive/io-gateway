#!/bin/bash
set -e
cd $(dirname "${BASH_SOURCE[0]}")
VERSION=${1:?version}
VERSION=${VERSION/-/}
mkdir -p deb/iogw_${VERSION}/usr/local/bin/
mkdir -p deb/iogw_${VERSION}/DEBIAN/
mkdir -p rpm
sed -e "/Version/cVersion: ${VERSION}" control > deb/iogw_${VERSION}/DEBIAN/control

cp bin/* deb/iogw_${VERSION}/usr/local/bin/

pushd deb
dpkg-deb --build iogw_${VERSION}
popd

pushd rpm
alien -rg --bump 0 ../deb/iogw_${VERSION}.deb
sed -i '/%dir/d' iogw-${VERSION}/iogw-${VERSION}-1.spec
cd iogw-${VERSION}
rpmbuild --buildroot $PWD -bb iogw-${VERSION}-1.spec
popd

