package l9plugins

import (
	"github.com/LeakIX/l9format"
	"github.com/xsssends/l9plugins/tcp"
	"github.com/xsssends/l9plugins/web"
)

var tcpPlugins = []l9format.ServicePluginInterface{
	tcp.Apache2449TraversalPlugin{},
	tcp.CouchDbOpenPlugin{},
	tcp.ElasticSearchExplorePlugin{},
	tcp.ElasticSearchOpenPlugin{},
	tcp.KafkaOpenPlugin{},
	tcp.MongoSchemaPlugin{},
	tcp.MongoOpenPlugin{},
	tcp.MysqlWeakPlugin{},
	tcp.MysqlSchemaPlugin{},
	tcp.RedisOpenPlugin{},
	tcp.SSHOpenPlugin{},
	tcp.DotDsStoreOpenPlugin{},
}

var webPlugins = []l9format.WebPluginInterface{
	web.ApacheStatusHttpPlugin{},
	web.ConfigJsonHttp{},
	web.DotEnvHttpPlugin{},
	web.GitConfigHttpPlugin{},
	web.IdxConfigPlugin{},
	web.LaravelTelescopeHttpPlugin{},
	web.PhpInfoHttpPlugin{},
	web.FirebaseHttpPlugin{},
	web.WpUserEnumHttp{},
	web.ConfluenceVersionIssue{},
	web.JiraPlugin{},
}

func GetTcpPlugins() []l9format.ServicePluginInterface {
	return tcpPlugins
}

func GetWebPlugins() []l9format.WebPluginInterface {
	return webPlugins
}
