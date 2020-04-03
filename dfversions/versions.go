package dfversions

const (
    max_uint32 = ^uint32(0)
    min_uint32 = uint32(0)
)

var (
    min_version = max_uint32
    max_version = min_uint32
)

var save_versions = map[uint32]string{
    1107: "0.21.93.19a",
    1108: "0.21.95.19b",
    1110: "0.21.100.19a",
    1113: "0.21.101.19a",
    1114: "0.21.101.19d",
    1117: "0.21.102.19a",
    1119: "0.21.104.19b",
    1121: "0.21.104.19d",
    1123: "0.21.104.21a",
    1125: "0.21.104.21b",
    1128: "0.22.107.21a",
    1131: "0.22.110.22b",
    1133: "0.22.110.22d",
    1134: "0.22.110.22e",
    1137: "0.22.110.22f",
    1148: "0.22.110.23c",
    1151: "0.22.120.23a",
    1159: "0.22.120.23b",
    1160: "0.22.121.23a",
    1161: "0.22.121.23b",
    1165: "0.22.123.23a",
    1167: "0.23.125.23a",
    1168: "0.23.125.23b",
    1169: "0.23.130.23a",
    1205: "0.27.169.32a",
    1206: "0.27.169.33a",
    1209: "0.27.169.33b",
    1211: "0.27.169.33c",
    1212: "0.27.169.33d",
    1213: "0.27.169.33e",
    1215: "0.27.169.33f",
    1216: "0.27.169.33g",
    1223: "0.27.173.38a",
    1231: "0.27.176.38a",
    1234: "0.27.176.38b",
    1235: "0.27.176.38c",
    1254: "0.28.181.39a",
    1255: "0.28.181.39b",
    1256: "0.28.181.39c",
    1259: "0.28.181.39d",
    1260: "0.28.181.39e",
    1261: "0.28.181.39f",
    1265: "0.28.181.40a",
    1266: "0.28.181.40b",
    1267: "0.28.181.40c",
    1268: "0.28.181.40d",
    1287: "0.31.01",
    1288: "0.31.02",
    1289: "0.31.03",
    1292: "0.31.04",
    1295: "0.31.05",
    1297: "0.31.06",
    1300: "0.31.08",
    1304: "0.31.09",
    1305: "0.31.10",
    1310: "0.31.11",
    1311: "0.31.12",
    1323: "0.31.13",
    1325: "0.31.14",
    1326: "0.31.15",
    1327: "0.31.16",
    1340: "0.31.17",
    1341: "0.31.18",
    1351: "0.31.19",
    1353: "0.31.20",
    1354: "0.31.21",
    1359: "0.31.22",
    1360: "0.31.23",
    1361: "0.31.24",
    1362: "0.31.25",
    1372: "0.34.01",
    1374: "0.34.02",
    1376: "0.34.03",
    1377: "0.34.04",
    1378: "0.34.05",
    1382: "0.34.06",
    1383: "0.34.07",
    1400: "0.34.08",
    1402: "0.34.09",
    1403: "0.34.10",
    1404: "0.34.11",
    1441: "0.40.01",
    1442: "0.40.02",
    1443: "0.40.03",
    1444: "0.40.04",
    1445: "0.40.05",
    1446: "0.40.06",
    1448: "0.40.07",
    1449: "0.40.08",
    1451: "0.40.09",
    1452: "0.40.10",
    1456: "0.40.11",
    1459: "0.40.12",
    1462: "0.40.13",
    1469: "0.40.14",
    1470: "0.40.15",
    1471: "0.40.16",
    1472: "0.40.17",
    1473: "0.40.18",
    1474: "0.40.19",
    1477: "0.40.20",
    1478: "0.40.21",
    1479: "0.40.22",
    1480: "0.40.23",
    1481: "0.40.24",
    1531: "0.42.01",
    1532: "0.42.02",
    1533: "0.42.03",
    1534: "0.42.04",
    1537: "0.42.05",
    1542: "0.42.06",
    1551: "0.43.01",
    1552: "0.43.02",
    1553: "0.43.03",
    1555: "0.43.04",
    1556: "0.43.05",
    1596: "0.44.01",
    1597: "0.44.02",
    1600: "0.44.03",
    1603: "0.44.04",
    1604: "0.44.05",
    1611: "0.44.06",
    1612: "0.44.07",
    1613: "0.44.08",
    1614: "0.44.09",
    1620: "0.44.10",
    1623: "0.44.11",
    1625: "0.44.12",
    1710: "0.47.01",
    1711: "0.47.02",
    1713: "0.47.03",
    1715: "0.47.04",
}

func init() {
    for id, _ := range save_versions {
        if id < min_version {
            min_version = id
        }
        if id > max_version {
            max_version = id
        }
    }
}

func IsKnown(id uint32) (ok bool) {
    _, ok = save_versions[id]
    return
}

func Describe(id uint32) string {
    if IsKnown(id) {
        return save_versions[id]
    } else if id < min_version {
        return "before " + save_versions[min_version]
    } else if id > max_version {
        return "after " + save_versions[max_version]
    } else {
        var prev, next uint32
        for prev = id; !IsKnown(prev); prev-- {}
        for next = id; !IsKnown(next); next++ {}
        return "between " + Describe(prev) + " and " + Describe(next)
    }
}
