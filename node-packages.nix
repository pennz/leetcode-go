# This file has been generated by node2nix 1.8.0. Do not edit!

{nodeEnv, fetchurl, fetchgit, globalBuildInputs ? []}:

let
  sources = {};
  args = {
    name = "leetcode-go";
    packageName = "leetcode-go";
    version = "1.0.0";
    src = ./.;
    buildInputs = globalBuildInputs;
    meta = {
      description = "[![Build Status](https://travis-ci.org/pennz/leetcode-go.svg?branch=master)](https://travis-ci.org/pennz/leetcode-go) [![Maintainability](https://api.codeclimate.com/v1/badges/c985a211d6be32adf6a2/maintainability)](https://codeclimate.com/github/pennz/leetcode-go/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/c985a211d6be32adf6a2/test_coverage)](https://codeclimate.com/github/pennz/leetcode-go/test_coverage)";
      homepage = "https://github.com/pennz/leetcode-go#readme";
      license = "ISC";
    };
    production = true;
    bypassCache = true;
    reconstructLock = false;
  };
in
{
  args = args;
  sources = sources;
  tarball = nodeEnv.buildNodeSourceDist args;
  package = nodeEnv.buildNodePackage args;
  shell = nodeEnv.buildNodeShell args;
}