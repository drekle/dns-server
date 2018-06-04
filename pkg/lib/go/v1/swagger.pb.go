package v1 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "api.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/record": {
      "get": {
        "operationId": "GetRecord",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1ResourceRecords"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "typeName",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "DNSService"
        ]
      },
      "delete": {
        "operationId": "DeleteRecord",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1ResourceRecord"
            }
          }
        },
        "parameters": [
          {
            "name": "header.name",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "header.typeName",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "a.address",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "aaaa.address",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "afsdb.subtype",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "afsdb.hostname",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "avc.txt",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          {
            "name": "caa.flag",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "caa.tag",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "caa.value",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "cdnskey.flags",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "cdnskey.protocol",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "cdnskey.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "cdnskey.publicKey",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "cds.keyTag",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "cds.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "cds.digestType",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "cds.digest",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "cert.keyTag",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "cert.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "cert.certificate",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "cname.target",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "csync.serial",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "csync.flags",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "csync.typeBitMap",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "integer",
              "format": "int64"
            }
          },
          {
            "name": "dhcid.digest",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "dlv.keyTag",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "dlv.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "dlv.digestType",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "dlv.digest",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "dname.target",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "dnskey.flags",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "dnskey.protocol",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "dnskey.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "dnskey.publicKey",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "ds.keyTag",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "ds.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "ds.digestType",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "ds.digest",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "eid.endpoint",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "eui48.address",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "uint64"
          },
          {
            "name": "eui64.address",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "uint64"
          },
          {
            "name": "gid.gid",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "gpos.latitude",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "gpos.longitude",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "gpos.altitude",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "hinfo.cpu",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "hinfo.os",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "hip.hitLength",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "hip.publicKeyAlgorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "hip.PublicKeyLength",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "hip.hit",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "hip.publicKey",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "hip.RendezvousServers",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "key.flags",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "key.protocol",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "key.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "key.publicKey",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "kx.preference",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "kx.exchanger",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "l32.preference",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "l32.Locator32",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "l64.preference",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "l64.Locator64",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "loc.Version",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "loc.Size",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "loc.HorizPre",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "loc.VertPre",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "loc.Latitude",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "loc.Longitude",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "loc.Altitude",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "lp.FQDN",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "mb.mb",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "md.md",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "mf.mf",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "mg.mg",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "minfo.rmail",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "minfo.email",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "mr.mr",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "mx.preference",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "mx.mx",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "naptr.order",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "naptr.preference",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "naptr.flags",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "naptr.service",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "naptr.regexp",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "naptr.replacement",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "nid.preference",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "nid.nodeid",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "uint64"
          },
          {
            "name": "nimloc.locator",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "ninfo.ZSDATA",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          {
            "name": "ns.ns",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "nsapptr.ptr",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "nsec.nextDomain",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "nsec.typeBitMap",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "integer",
              "format": "int64"
            }
          },
          {
            "name": "nsec3.hash",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "nsec3.flags",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "nsec3.iterations",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "nsec3.saltLength",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "nsec3.salt",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "nsec3.hashLength",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "nsec3.nextDomain",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "nsec3.typeBitMap",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "integer",
              "format": "int64"
            }
          },
          {
            "name": "nsec3param.hash",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "nsec3param.flags",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "nsec3param.iterations",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "nsec3param.saltLength",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "nsec3param.salt",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "openpgpkey.publicKey",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "ptr.prt",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "px.preference",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "px.map822",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "px.mapx400",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "rkey.flags",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rkey.protocol",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rkey.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rkey.publicKey",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "rp.mbox",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "rp.txt",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "rrsig.typeCovered",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rrsig.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rrsig.labels",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rrsig.origttl",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rrsig.expiration",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rrsig.inception",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rrsig.keytag",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rrsig.signerName",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "rrsig.Signature",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "rt.preference",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "rt.host",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "sig.typeCovered",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "sig.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "sig.labels",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "sig.origttl",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "sig.expiration",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "sig.inception",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "sig.keytag",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "sig.signerName",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "sig.Signature",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "smimea.usage",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "smimea.selector",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "smimea.matchingType",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "smimea.certificate",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "soa.ns",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "soa.mbox",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "soa.serial",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "soa.refresh",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "soa.retry",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "soa.expire",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "soa.minttl",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "spf.txt",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          {
            "name": "srv.priority",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "srv.weight",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "srv.port",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "srv.target",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "sshfp.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "sshfp.type",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "sshfp.fingerprint",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "ta.keyTag",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "ta.algorithm",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "ta.digestType",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "ta.digest",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "talink.previousName",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "talink.nextName",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "tkey.algorithm",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "tkey.inception",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tkey.expiration",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tkey.mode",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tkey.error",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tkey.keySize",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tkey.key",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "tkey.otherLen",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tkey.otherData",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "tlsa.usage",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tlsa.selector",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tlsa.matchingType",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tlsa.certificate",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "tsig.algorithm",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "tsig.timeSigned",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "uint64"
          },
          {
            "name": "tsig.fudge",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tsig.MACSize",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tsig.MAC",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "tsig.origID",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tsig.error",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tsig.otherLen",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "tsig.otherData",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "txt.txt",
            "in": "query",
            "required": false,
            "type": "array",
            "items": {
              "type": "string"
            }
          },
          {
            "name": "uid.uid",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "uinfo.uinfo",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "uri.priority",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "uri.weight",
            "in": "query",
            "required": false,
            "type": "integer",
            "format": "int64"
          },
          {
            "name": "uri.target",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "x25.PSDNAddress",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "DNSService"
        ]
      },
      "post": {
        "operationId": "PostRecord",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1ResourceRecord"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1ResourceRecord"
            }
          }
        ],
        "tags": [
          "DNSService"
        ]
      },
      "put": {
        "operationId": "PutRecord",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v1ResourceRecord"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1ResourceRecord"
            }
          }
        ],
        "tags": [
          "DNSService"
        ]
      }
    }
  },
  "definitions": {
    "v1A": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        }
      }
    },
    "v1AAAA": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        }
      }
    },
    "v1AFSDB": {
      "type": "object",
      "properties": {
        "subtype": {
          "type": "integer",
          "format": "int64"
        },
        "hostname": {
          "type": "string"
        }
      }
    },
    "v1AVC": {
      "type": "object",
      "properties": {
        "txt": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v1CAA": {
      "type": "object",
      "properties": {
        "flag": {
          "type": "integer",
          "format": "int64"
        },
        "tag": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "v1CDNSKEY": {
      "type": "object",
      "properties": {
        "flags": {
          "type": "integer",
          "format": "int64"
        },
        "protocol": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "publicKey": {
          "type": "string"
        }
      }
    },
    "v1CDS": {
      "type": "object",
      "properties": {
        "keyTag": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "digestType": {
          "type": "integer",
          "format": "int64"
        },
        "digest": {
          "type": "string"
        }
      }
    },
    "v1CERT": {
      "type": "object",
      "properties": {
        "keyTag": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "certificate": {
          "type": "string"
        }
      }
    },
    "v1CNAME": {
      "type": "object",
      "properties": {
        "target": {
          "type": "string"
        }
      }
    },
    "v1CSYNC": {
      "type": "object",
      "properties": {
        "serial": {
          "type": "integer",
          "format": "int64"
        },
        "flags": {
          "type": "integer",
          "format": "int64"
        },
        "typeBitMap": {
          "type": "array",
          "items": {
            "type": "integer",
            "format": "int64"
          }
        }
      }
    },
    "v1DHCID": {
      "type": "object",
      "properties": {
        "digest": {
          "type": "string"
        }
      }
    },
    "v1DLV": {
      "type": "object",
      "properties": {
        "keyTag": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "digestType": {
          "type": "integer",
          "format": "int64"
        },
        "digest": {
          "type": "string"
        }
      }
    },
    "v1DNAME": {
      "type": "object",
      "properties": {
        "target": {
          "type": "string"
        }
      }
    },
    "v1DNSKEY": {
      "type": "object",
      "properties": {
        "flags": {
          "type": "integer",
          "format": "int64"
        },
        "protocol": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "publicKey": {
          "type": "string"
        }
      }
    },
    "v1DS": {
      "type": "object",
      "properties": {
        "keyTag": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "digestType": {
          "type": "integer",
          "format": "int64"
        },
        "digest": {
          "type": "string"
        }
      }
    },
    "v1EID": {
      "type": "object",
      "properties": {
        "endpoint": {
          "type": "string"
        }
      }
    },
    "v1EUI48": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "v1EUI64": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "v1GID": {
      "type": "object",
      "properties": {
        "gid": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "v1GPOS": {
      "type": "object",
      "properties": {
        "latitude": {
          "type": "string"
        },
        "longitude": {
          "type": "string"
        },
        "altitude": {
          "type": "string"
        }
      }
    },
    "v1HINFO": {
      "type": "object",
      "properties": {
        "cpu": {
          "type": "string"
        },
        "os": {
          "type": "string"
        }
      }
    },
    "v1HIP": {
      "type": "object",
      "properties": {
        "hitLength": {
          "type": "integer",
          "format": "int64"
        },
        "publicKeyAlgorithm": {
          "type": "integer",
          "format": "int64"
        },
        "PublicKeyLength": {
          "type": "integer",
          "format": "int64"
        },
        "hit": {
          "type": "string"
        },
        "publicKey": {
          "type": "string"
        },
        "RendezvousServers": {
          "type": "string"
        }
      }
    },
    "v1Header": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "typeName": {
          "type": "string"
        }
      }
    },
    "v1KEY": {
      "type": "object",
      "properties": {
        "flags": {
          "type": "integer",
          "format": "int64"
        },
        "protocol": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "publicKey": {
          "type": "string"
        }
      }
    },
    "v1KX": {
      "type": "object",
      "properties": {
        "preference": {
          "type": "integer",
          "format": "int64"
        },
        "exchanger": {
          "type": "string"
        }
      }
    },
    "v1L32": {
      "type": "object",
      "properties": {
        "preference": {
          "type": "integer",
          "format": "int64"
        },
        "Locator32": {
          "type": "string"
        }
      }
    },
    "v1L64": {
      "type": "object",
      "properties": {
        "preference": {
          "type": "integer",
          "format": "int64"
        },
        "Locator64": {
          "type": "string"
        }
      }
    },
    "v1LOC": {
      "type": "object",
      "properties": {
        "Version": {
          "type": "integer",
          "format": "int64"
        },
        "Size": {
          "type": "integer",
          "format": "int64"
        },
        "HorizPre": {
          "type": "integer",
          "format": "int64"
        },
        "VertPre": {
          "type": "integer",
          "format": "int64"
        },
        "Latitude": {
          "type": "integer",
          "format": "int64"
        },
        "Longitude": {
          "type": "integer",
          "format": "int64"
        },
        "Altitude": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "v1LP": {
      "type": "object",
      "properties": {
        "FQDN": {
          "type": "string"
        }
      }
    },
    "v1MB": {
      "type": "object",
      "properties": {
        "mb": {
          "type": "string"
        }
      }
    },
    "v1MD": {
      "type": "object",
      "properties": {
        "md": {
          "type": "string"
        }
      }
    },
    "v1MF": {
      "type": "object",
      "properties": {
        "mf": {
          "type": "string"
        }
      }
    },
    "v1MG": {
      "type": "object",
      "properties": {
        "mg": {
          "type": "string"
        }
      }
    },
    "v1MINFO": {
      "type": "object",
      "properties": {
        "rmail": {
          "type": "string"
        },
        "email": {
          "type": "string"
        }
      }
    },
    "v1MR": {
      "type": "object",
      "properties": {
        "mr": {
          "type": "string"
        }
      }
    },
    "v1MX": {
      "type": "object",
      "properties": {
        "preference": {
          "type": "integer",
          "format": "int64"
        },
        "mx": {
          "type": "string"
        }
      }
    },
    "v1NAPTR": {
      "type": "object",
      "properties": {
        "order": {
          "type": "integer",
          "format": "int64"
        },
        "preference": {
          "type": "integer",
          "format": "int64"
        },
        "flags": {
          "type": "string"
        },
        "service": {
          "type": "string"
        },
        "regexp": {
          "type": "string"
        },
        "replacement": {
          "type": "string"
        }
      }
    },
    "v1NID": {
      "type": "object",
      "properties": {
        "preference": {
          "type": "integer",
          "format": "int64"
        },
        "nodeid": {
          "type": "string",
          "format": "uint64"
        }
      }
    },
    "v1NIMLOC": {
      "type": "object",
      "properties": {
        "locator": {
          "type": "string"
        }
      }
    },
    "v1NINFO": {
      "type": "object",
      "properties": {
        "ZSDATA": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v1NS": {
      "type": "object",
      "properties": {
        "ns": {
          "type": "string"
        }
      }
    },
    "v1NSAPPTR": {
      "type": "object",
      "properties": {
        "ptr": {
          "type": "string"
        }
      }
    },
    "v1NSEC": {
      "type": "object",
      "properties": {
        "nextDomain": {
          "type": "string"
        },
        "typeBitMap": {
          "type": "array",
          "items": {
            "type": "integer",
            "format": "int64"
          }
        }
      }
    },
    "v1NSEC3": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "integer",
          "format": "int64"
        },
        "flags": {
          "type": "integer",
          "format": "int64"
        },
        "iterations": {
          "type": "integer",
          "format": "int64"
        },
        "saltLength": {
          "type": "integer",
          "format": "int64"
        },
        "salt": {
          "type": "string"
        },
        "hashLength": {
          "type": "integer",
          "format": "int64"
        },
        "nextDomain": {
          "type": "string"
        },
        "typeBitMap": {
          "type": "array",
          "items": {
            "type": "integer",
            "format": "int64"
          }
        }
      }
    },
    "v1NSEC3PARAM": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "integer",
          "format": "int64"
        },
        "flags": {
          "type": "integer",
          "format": "int64"
        },
        "iterations": {
          "type": "integer",
          "format": "int64"
        },
        "saltLength": {
          "type": "integer",
          "format": "int64"
        },
        "salt": {
          "type": "string"
        }
      }
    },
    "v1OPENPGPKEY": {
      "type": "object",
      "properties": {
        "publicKey": {
          "type": "string"
        }
      }
    },
    "v1OPT": {
      "type": "object",
      "title": "TODO:"
    },
    "v1PTR": {
      "type": "object",
      "properties": {
        "prt": {
          "type": "string"
        }
      }
    },
    "v1PX": {
      "type": "object",
      "properties": {
        "preference": {
          "type": "integer",
          "format": "int64"
        },
        "map822": {
          "type": "string"
        },
        "mapx400": {
          "type": "string"
        }
      }
    },
    "v1RKEY": {
      "type": "object",
      "properties": {
        "flags": {
          "type": "integer",
          "format": "int64"
        },
        "protocol": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "publicKey": {
          "type": "string"
        }
      }
    },
    "v1RP": {
      "type": "object",
      "properties": {
        "mbox": {
          "type": "string"
        },
        "txt": {
          "type": "string"
        }
      }
    },
    "v1RRSIG": {
      "type": "object",
      "properties": {
        "typeCovered": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "labels": {
          "type": "integer",
          "format": "int64"
        },
        "origttl": {
          "type": "integer",
          "format": "int64"
        },
        "expiration": {
          "type": "integer",
          "format": "int64"
        },
        "inception": {
          "type": "integer",
          "format": "int64"
        },
        "keytag": {
          "type": "integer",
          "format": "int64"
        },
        "signerName": {
          "type": "string"
        },
        "Signature": {
          "type": "string"
        }
      }
    },
    "v1RT": {
      "type": "object",
      "properties": {
        "preference": {
          "type": "integer",
          "format": "int64"
        },
        "host": {
          "type": "string"
        }
      }
    },
    "v1ResourceRecord": {
      "type": "object",
      "properties": {
        "header": {
          "$ref": "#/definitions/v1Header"
        },
        "a": {
          "$ref": "#/definitions/v1A"
        },
        "aaaa": {
          "$ref": "#/definitions/v1AAAA"
        },
        "afsdb": {
          "$ref": "#/definitions/v1AFSDB"
        },
        "avc": {
          "$ref": "#/definitions/v1AVC"
        },
        "caa": {
          "$ref": "#/definitions/v1CAA"
        },
        "cdnskey": {
          "$ref": "#/definitions/v1CDNSKEY"
        },
        "cds": {
          "$ref": "#/definitions/v1CDS"
        },
        "cert": {
          "$ref": "#/definitions/v1CERT"
        },
        "cname": {
          "$ref": "#/definitions/v1CNAME"
        },
        "csync": {
          "$ref": "#/definitions/v1CSYNC"
        },
        "dhcid": {
          "$ref": "#/definitions/v1DHCID"
        },
        "dlv": {
          "$ref": "#/definitions/v1DLV"
        },
        "dname": {
          "$ref": "#/definitions/v1DNAME"
        },
        "dnskey": {
          "$ref": "#/definitions/v1DNSKEY"
        },
        "ds": {
          "$ref": "#/definitions/v1DS"
        },
        "eid": {
          "$ref": "#/definitions/v1EID"
        },
        "eui48": {
          "$ref": "#/definitions/v1EUI48"
        },
        "eui64": {
          "$ref": "#/definitions/v1EUI64"
        },
        "gid": {
          "$ref": "#/definitions/v1GID"
        },
        "gpos": {
          "$ref": "#/definitions/v1GPOS"
        },
        "hinfo": {
          "$ref": "#/definitions/v1HINFO"
        },
        "hip": {
          "$ref": "#/definitions/v1HIP"
        },
        "key": {
          "$ref": "#/definitions/v1KEY"
        },
        "kx": {
          "$ref": "#/definitions/v1KX"
        },
        "l32": {
          "$ref": "#/definitions/v1L32"
        },
        "l64": {
          "$ref": "#/definitions/v1L64"
        },
        "loc": {
          "$ref": "#/definitions/v1LOC"
        },
        "lp": {
          "$ref": "#/definitions/v1LP"
        },
        "mb": {
          "$ref": "#/definitions/v1MB"
        },
        "md": {
          "$ref": "#/definitions/v1MD"
        },
        "mf": {
          "$ref": "#/definitions/v1MF"
        },
        "mg": {
          "$ref": "#/definitions/v1MG"
        },
        "minfo": {
          "$ref": "#/definitions/v1MINFO"
        },
        "mr": {
          "$ref": "#/definitions/v1MR"
        },
        "mx": {
          "$ref": "#/definitions/v1MX"
        },
        "naptr": {
          "$ref": "#/definitions/v1NAPTR"
        },
        "nid": {
          "$ref": "#/definitions/v1NID"
        },
        "nimloc": {
          "$ref": "#/definitions/v1NIMLOC"
        },
        "ninfo": {
          "$ref": "#/definitions/v1NINFO"
        },
        "ns": {
          "$ref": "#/definitions/v1NS"
        },
        "nsapptr": {
          "$ref": "#/definitions/v1NSAPPTR"
        },
        "nsec": {
          "$ref": "#/definitions/v1NSEC"
        },
        "nsec3": {
          "$ref": "#/definitions/v1NSEC3"
        },
        "nsec3param": {
          "$ref": "#/definitions/v1NSEC3PARAM"
        },
        "openpgpkey": {
          "$ref": "#/definitions/v1OPENPGPKEY"
        },
        "opt": {
          "$ref": "#/definitions/v1OPT"
        },
        "ptr": {
          "$ref": "#/definitions/v1PTR"
        },
        "px": {
          "$ref": "#/definitions/v1PX"
        },
        "rkey": {
          "$ref": "#/definitions/v1RKEY"
        },
        "rp": {
          "$ref": "#/definitions/v1RP"
        },
        "rrsig": {
          "$ref": "#/definitions/v1RRSIG"
        },
        "rt": {
          "$ref": "#/definitions/v1RT"
        },
        "sig": {
          "$ref": "#/definitions/v1SIG"
        },
        "smimea": {
          "$ref": "#/definitions/v1SMIMEA"
        },
        "soa": {
          "$ref": "#/definitions/v1SOA"
        },
        "spf": {
          "$ref": "#/definitions/v1SPF"
        },
        "srv": {
          "$ref": "#/definitions/v1SRV"
        },
        "sshfp": {
          "$ref": "#/definitions/v1SSHFP"
        },
        "ta": {
          "$ref": "#/definitions/v1TA"
        },
        "talink": {
          "$ref": "#/definitions/v1TALINK"
        },
        "tkey": {
          "$ref": "#/definitions/v1TKEY"
        },
        "tlsa": {
          "$ref": "#/definitions/v1TLSA"
        },
        "tsig": {
          "$ref": "#/definitions/v1TSIG"
        },
        "txt": {
          "$ref": "#/definitions/v1TXT"
        },
        "uid": {
          "$ref": "#/definitions/v1UID"
        },
        "uinfo": {
          "$ref": "#/definitions/v1UINFO"
        },
        "uri": {
          "$ref": "#/definitions/v1URI"
        },
        "x25": {
          "$ref": "#/definitions/v1X25"
        }
      }
    },
    "v1ResourceRecords": {
      "type": "object",
      "properties": {
        "records": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1ResourceRecord"
          }
        }
      }
    },
    "v1SIG": {
      "type": "object",
      "properties": {
        "typeCovered": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "labels": {
          "type": "integer",
          "format": "int64"
        },
        "origttl": {
          "type": "integer",
          "format": "int64"
        },
        "expiration": {
          "type": "integer",
          "format": "int64"
        },
        "inception": {
          "type": "integer",
          "format": "int64"
        },
        "keytag": {
          "type": "integer",
          "format": "int64"
        },
        "signerName": {
          "type": "string"
        },
        "Signature": {
          "type": "string"
        }
      }
    },
    "v1SMIMEA": {
      "type": "object",
      "properties": {
        "usage": {
          "type": "integer",
          "format": "int64"
        },
        "selector": {
          "type": "integer",
          "format": "int64"
        },
        "matchingType": {
          "type": "integer",
          "format": "int64"
        },
        "certificate": {
          "type": "string"
        }
      }
    },
    "v1SOA": {
      "type": "object",
      "properties": {
        "ns": {
          "type": "string"
        },
        "mbox": {
          "type": "string"
        },
        "serial": {
          "type": "integer",
          "format": "int64"
        },
        "refresh": {
          "type": "integer",
          "format": "int64"
        },
        "retry": {
          "type": "integer",
          "format": "int64"
        },
        "expire": {
          "type": "integer",
          "format": "int64"
        },
        "minttl": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "v1SPF": {
      "type": "object",
      "properties": {
        "txt": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v1SRV": {
      "type": "object",
      "properties": {
        "priority": {
          "type": "integer",
          "format": "int64"
        },
        "weight": {
          "type": "integer",
          "format": "int64"
        },
        "port": {
          "type": "integer",
          "format": "int64"
        },
        "target": {
          "type": "string"
        }
      }
    },
    "v1SSHFP": {
      "type": "object",
      "properties": {
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "type": {
          "type": "integer",
          "format": "int64"
        },
        "fingerprint": {
          "type": "string"
        }
      }
    },
    "v1TA": {
      "type": "object",
      "properties": {
        "keyTag": {
          "type": "integer",
          "format": "int64"
        },
        "algorithm": {
          "type": "integer",
          "format": "int64"
        },
        "digestType": {
          "type": "integer",
          "format": "int64"
        },
        "digest": {
          "type": "string"
        }
      }
    },
    "v1TALINK": {
      "type": "object",
      "properties": {
        "previousName": {
          "type": "string"
        },
        "nextName": {
          "type": "string"
        }
      }
    },
    "v1TKEY": {
      "type": "object",
      "properties": {
        "algorithm": {
          "type": "string"
        },
        "inception": {
          "type": "integer",
          "format": "int64"
        },
        "expiration": {
          "type": "integer",
          "format": "int64"
        },
        "mode": {
          "type": "integer",
          "format": "int64"
        },
        "error": {
          "type": "integer",
          "format": "int64"
        },
        "keySize": {
          "type": "integer",
          "format": "int64"
        },
        "key": {
          "type": "string"
        },
        "otherLen": {
          "type": "integer",
          "format": "int64"
        },
        "otherData": {
          "type": "string"
        }
      }
    },
    "v1TLSA": {
      "type": "object",
      "properties": {
        "usage": {
          "type": "integer",
          "format": "int64"
        },
        "selector": {
          "type": "integer",
          "format": "int64"
        },
        "matchingType": {
          "type": "integer",
          "format": "int64"
        },
        "certificate": {
          "type": "string"
        }
      }
    },
    "v1TSIG": {
      "type": "object",
      "properties": {
        "algorithm": {
          "type": "string"
        },
        "timeSigned": {
          "type": "string",
          "format": "uint64"
        },
        "fudge": {
          "type": "integer",
          "format": "int64"
        },
        "MACSize": {
          "type": "integer",
          "format": "int64"
        },
        "MAC": {
          "type": "string"
        },
        "origID": {
          "type": "integer",
          "format": "int64"
        },
        "error": {
          "type": "integer",
          "format": "int64"
        },
        "otherLen": {
          "type": "integer",
          "format": "int64"
        },
        "otherData": {
          "type": "string"
        }
      }
    },
    "v1TXT": {
      "type": "object",
      "properties": {
        "txt": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "v1UID": {
      "type": "object",
      "properties": {
        "uid": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "v1UINFO": {
      "type": "object",
      "properties": {
        "uinfo": {
          "type": "string"
        }
      }
    },
    "v1URI": {
      "type": "object",
      "properties": {
        "priority": {
          "type": "integer",
          "format": "int64"
        },
        "weight": {
          "type": "integer",
          "format": "int64"
        },
        "target": {
          "type": "string"
        }
      }
    },
    "v1X25": {
      "type": "object",
      "properties": {
        "PSDNAddress": {
          "type": "string"
        }
      }
    }
  }
}
`
)
