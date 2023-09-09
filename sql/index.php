<?php

date_default_timezone_set("Asia/Shanghai");
error_reporting(E_ERROR);

//365才执行
$exchange = @$_GET['exchange'] ;
$doadd = @$_GET['doadd'] ;
$dodel = 0 ;

//待同步数据库--正式库--一般设置localhost
$selfConf = array( 
    'host'  => 'localhost',
    'user'  => 'user',
    'pwd'   => '123456',
    'db'    => 'go_admin'
);
//同步来源数据库 -- 设置为测试数据库地址
$sourceConf = array( 
    'host'  => 'localhost',
    'user'  => '123456',
    'pwd'   => 'root',
    'db'    => 'goadmin'
);

//交换
if($exchange == 66){
    $temp = $selfConf;
    $selfConf = $sourceConf ;
    $sourceConf = $temp ;
}

?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>数据库同步</title>
    <!-- Bootstrap -->
    <link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
    <!-- jQuery (Bootstrap 的所有 JavaScript 插件都依赖 jQuery，所以必须放在前边) -->
    <script src="https://cdn.bootcss.com/jquery/1.12.4/jquery.min.js"></script>
    <!-- 加载 Bootstrap 的所有 JavaScript 插件。你也可以根据需要只加载单个插件。 -->
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
    
    <style type="text/css">
        table.gridtable {
            font-family: "微软雅黑";
            font-size: 12px;
            width: 100%;
            color: #333333;
            border-width: 1px;
            border-color: #666666;
            border-collapse: collapse;
        }

        table.gridtable th {
            border-width: 1px;
            padding: 10px;
            border-style: solid;
            border-color: #666666;
            background-color: #dedede;
        }

        table.gridtable td {
            border-width: 1px;
            padding: 10px;
            border-style: solid;
            border-color: #666666;
            background-color: #ffffff;
        }
    </style>
    
</head>
<body >
<div style="text-align: center; color: #00aa00;">
	<h1>快租365商城项目：<?php echo $selfConf['db'];?>同步源库<?php echo $sourceConf['db'];?></h1>
</div>

<div class="col-lg-12 col-sm-12 col-xs-12" >
    <div class="alert alert-danger" role="alert" >
    <strong>Tips!</strong>
      URL参数中  doadd=365  时，执行对应add 操作 ,  exchange=66 时,数据库配置交换。
    </div>
    
<?php

class MysqlSync{
    
    /**
     * 执行状态记录
     * @var array
     */
    private $stat = array();
    
    /**
     * 默认值需要加上引号的类型的索引
     * @var array
     */
    private $convert_map = array('varchar', 'char', 'tinytext', 'mediumtext', 'text', 'longtext', 'enum');
    
    /**
     * 数据库结构同步
     * @param $selfConf
     * @param $sourceConf
     * @return array
     */
    function sync($selfConf, $sourceConf,$doadd='0' , $dodel ="0")
    {   
        $change= [] ;
        $self = new \Mysqli($selfConf['host'], $selfConf['user'], $selfConf['pwd'], $selfConf['db']);
        $source = new \Mysqli($sourceConf['host'], $sourceConf['user'], $sourceConf['pwd'], $sourceConf['db']);
        if (!$self) {
            die($selfConf['db']."连接失败: " . $self->connect_error);
        }   
        if (!$source) {
            die($sourceConf['db']."连接失败: " . $source->connect_error);
        }   
        $selfData = $this->getStructure($self, $selfConf['db']);     //获取本身，和对比源的结构
        $sourceData = $this->getStructure($source, $sourceConf['db']);
        
        $selfKeys = array_keys($selfData);      //获取本身，和对比源的表
        $sourceKeys = array_keys($sourceData);
        
        $removeList = array_diff($selfKeys, $sourceKeys);       //如果自身有，源没有，就删除
        $createList = array_diff($sourceKeys, $selfKeys);       //如果源有，自身没有，就新增
        
        if(!empty($removeList)){        //执行删除操作
            $remove_tab = '';
            foreach($removeList as $val){
                $remove_tab .= "`{$val}`,";
            }
            $remove_tab = trim($remove_tab, ',');
            $remove_sql = "DROP TABLE {$remove_tab}";
            
            $sqllog = array();
            $sqllog['db'] = $selfConf['host'].'  '.$selfConf['db'];
            $sqllog['table'] = $remove_tab;
            $sqllog['column'] = '';
            $sqllog['sql'] = $remove_sql;
            
            if($dodel == '365'){
                if($self->query($remove_sql)){
                    $sqllog['status'] = 1;
                }else{
                    $sqllog['status'] = 2;
                }
            }
            $this->stat['del'][] = $sqllog;
        }
        
        if(!empty($createList)){        //执行新增操作
            foreach($createList as $val){
                $create_arr = array();
                foreach($sourceData[$val] as $item){
                    $sql_write = "`{$item['COLUMN_NAME']}` {$item['COLUMN_TYPE']}";
                    if(!empty($item['COLUMN_DEFAULT'])){
                        if(in_array($item['DATA_TYPE'], $this->convert_map)){
                            $sql_write .= " DEFAULT '{$item['COLUMN_DEFAULT']}'";
                        }else{
                            $sql_write .= " DEFAULT {$item['COLUMN_DEFAULT']}";
                        }
                    }
                    $sql_write .= " COMMENT '{$item['COLUMN_COMMENT']} ' ";
                    $create_arr[] = $sql_write;
                }
                $create_sql = "CREATE TABLE IF NOT EXISTS `{$val}` (" . implode(',', $create_arr) . ")";
                
                $sqllog = array();
                $sqllog['db'] = $selfConf['host'].'  '.$selfConf['db'];
                $sqllog['table'] = $val;
                $sqllog['column'] = '';
                $sqllog['sql'] = $create_sql;
                
                if($doadd == '365'){
                    if($self->query($create_sql)){
                        $sqllog['status'] = 1;
                    }else{
                        $sqllog['status'] = 2;
                    }
                }
                $this->stat['add'][] = $sqllog;
                
            }
        }
        
        foreach($sourceData as $pKey => $item){     //对比表的字段是否相同
            foreach($selfData as $key => $val){
                if($pKey == $key){     //检测表结构是否相同
                    $removeColumn = array_diff_key($val, $item);
                    $addColumn = array_diff_key($item, $val);
                    if(!empty($removeColumn)){
                        foreach($removeColumn as $removeVal){
                            $removeColumnSql = "ALTER TABLE `{$key}` DROP COLUMN `{$removeVal['COLUMN_NAME']}`";
                            
                            $sqllog = array();
                            $sqllog['db'] = $selfConf['host'].'  '.$selfConf['db'];
                            $sqllog['table'] = $key;
                            $sqllog['column'] = $removeVal['COLUMN_NAME'];
                            $sqllog['sql'] = $removeColumnSql;
                            
                            if($dodel== '365'){
                                if($self->query($removeColumnSql)){
                                    $sqllog['status'] = 1;
                                }else{
                                    $sqllog['status'] = 2;
                                }
                            }
                            $this->stat['del'][] = $sqllog;
                        }
                    }
                    if(!empty($addColumn)){
                        foreach($addColumn as $addVal){
                            $addInfo = "`{$addVal['COLUMN_NAME']}` {$addVal['COLUMN_TYPE']}";
                            if(!empty($addVal['COLUMN_DEFAULT'])){
                                if(in_array($addVal['DATA_TYPE'], $this->convert_map)){
                                    $addInfo .= " DEFAULT '{$addVal['COLUMN_DEFAULT']}'";
                                }else{
                                    $addInfo .= " DEFAULT {$addVal['COLUMN_DEFAULT']}";
                                }
                            }
                            $addInfo .= " COMMENT '{$addVal['COLUMN_COMMENT']} ' ";
                            $addSql = "ALTER TABLE `{$key}` ADD COLUMN {$addInfo}";
                            
                            $sqllog = array();
                            $sqllog['db']     = $selfConf['host'].'  '.$selfConf['db'];
                            $sqllog['table']  = $key;
                            $sqllog['column'] = $addVal['COLUMN_NAME'];
                            $sqllog['sql']    = $addSql;
                            
                            if($doadd == '365'){
                                if($self->query($addSql)){
                                    $sqllog['status'] = 1;
                                }else{
                                    $sqllog['status'] = 2;
                                }
                            }
                            $this->stat['add'][] = $sqllog;
                            
                        }
                    }
                }
            }
        }
        $resultee['dbone'] = $change ;
        $resultee['dbone'] = array('conn'=>$self ,'conf'=> $selfConf);
        $resultee['dbtwo'] = array('conn'=>$source ,'conf'=> $sourceConf);
        $resultee['res'] = $this->stat;
        return $resultee;
    }
    
    /**
     * 获取表结构
     * @param $resource
     * @param $db
     * @return array
     */
    function getStructure($resource, $db){
        $table_str = '';
        $info = array();
        $sql_table = 'SHOW TABLES';
        $res_table = $resource->query($sql_table);
        while($row_table = $res_table->fetch_assoc()){
            $table_str .= "'" . current($row_table) . "',";
        }
        $table_str = trim($table_str, ',');
        $column_sql = "SELECT * FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name IN({$table_str}) AND table_schema = '{$db}'";
        $column_res = $resource->query($column_sql);
        if($column_res) {
            while ($row_column = $column_res->fetch_assoc()) {
                $info[] = $row_column;
            }
            return $this->gen($info);
        }else{
            return array();
        }
    }
    
    /**
     * 数据排序处理
     * @param $array
     * @return array
     */
    function gen($array){
        $data = array();
        foreach($array as $key => $item){
            if(!array_key_exists($item['TABLE_NAME'], $data)) {
                foreach ($array as $value) {
                    if ($value['TABLE_NAME'] == $item['TABLE_NAME']) {
                        $data[$item['TABLE_NAME']][$value['COLUMN_NAME']] = $value;
                    }
                }
            }
        }
        return $data;
    }
    
}
?>

<div class="alert alert-success" role="alert" >
<strong>Well done!</strong>
结构同步执行开始<?php echo date("Y-m-d H:i:s");  ?>
<?php 

$sync = new MysqlSync();
$result = $sync->sync($selfConf, $sourceConf,$doadd,$dodel);
$res = $result['res'] ;

?>
执行结束<?php echo date("Y-m-d H:i:s");  ?>
</div>
<div class="table-responsive">
    <table class="table table-bordered table-striped">
      <thead>
      	<tr><td colspan="6"><strong>需要新增字段</strong></td></tr>
        <tr>
          <td style="width:5%;" ></td>
          <td style="width:15%;">数据库</td>
          <th style="width:15%;">表格</th>
          <th style="width:15%;">字段</th>
          <th style="width:40%;">sql语句</th>
          <th style="width:10%;">执行状态</th>
        </tr>
      </thead>
      <tbody>
      <?php foreach($res['add'] as $k=>$val){ ?>
        <tr>
          <th scope="row"><?php echo $k+1; ?></th>
          <td class="text-danger"> <?php echo $val['db'] ;?> </td>
          <td class="text-success"> <?php echo $val['table'] ;?> </td>
          <td class="text-muted"> <?php echo $val['column'] ;?> </td>
          <td class="text-muted"> <?php echo $val['sql'] ;?> </td>
          <td class="text-muted">
          <?php if($val['status'] == 1){ ?>
          	<span class="glyphicon glyphicon-ok" aria-hidden="true"></span>执行成功
          	<?php }else if($val['status'] == 2){?>
          	<span class="glyphicon glyphicon-remove" aria-hidden="true"></span>执行失败
          	<?php }  ?>
          	</td>
        </tr>
       <?php } ?>
      </tbody>
      
       <thead>
        <tr><td colspan="6"><strong>需要删除字段</strong></td></tr>
        <tr>
            <td style="width:5%;" ></td>
            <td style="width:15%;">数据库</td>
            <th style="width:15%;">表格</th>
            <th style="width:15%;">字段</th>
            <th style="width:40%;">sql语句</th>
            <th style="width:10%;">执行状态</th>
        </tr>
      </thead>
      <tbody>
      <?php foreach($res['del'] as $k=>$val){ ?>
        <tr>
          <th scope="row"><?php echo $k+1; ?></th>
          <td class="text-danger"> <?php echo $val['db'] ;?> </td>
          <td class="text-success"> <?php echo $val['table'] ;?> </td>
          <td class="text-muted"> <?php echo $val['column'] ;?> </td>
          <td class="text-muted"> <?php echo $val['sql'] ;?> </td>
          <td class="text-muted">
          	<?php if($val['status'] == 1){ ?>
          	<span class="glyphicon glyphicon-ok" aria-hidden="true"></span>执行成功
          	<?php }else if($val['status'] == 2){ ?>
          	<span class="glyphicon glyphicon-remove" aria-hidden="true"></span>执行失败
          	<?php }  ?>
          </td>
        </tr>
         <?php } ?>
      </tbody>
      
    </table>
  </div>
  
<?php 

$dbone = $result['dbone']['conn'];
$dbone->query('set names utf8');
$sql_tables = "SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.`TABLES` WHERE TABLE_SCHEMA = '" . $selfConf['db'] . "'"; //查询表名及其备注
$ret_tables = $dbone->query($sql_tables);
if (!$ret_tables) {
    echo '未查询到表单';
    die;
}

?>

<div class="col-lg-6 col-md-6 col-sm-6 col-xs-6"  >
	<div class="page-header">
      <h1><?php  echo $selfConf['host']."  ".$selfConf['db'] ; ?></h1>
    </div>
<?php

$i = 1;
while ($row_tables = mysqli_fetch_assoc($ret_tables)) {
    $currentTable = $row_tables['TABLE_NAME'];// 表的名称
    ?>
<div class="panel panel-info">
      <div class="panel-heading">
        <h3 class="panel-title">表序号：<?php echo $i . PHP_EOL . "&nbsp;&nbsp;"; ?> 
        <b>表名：<?php echo $row_tables['TABLE_NAME'] ?></b>　　
        <b>备注：<?php echo $row_tables['TABLE_COMMENT'] ?></b>
        </h3>
      </div>
      <div class="panel-body">
    <?php
    $sql_column = "SELECT
				`COLUMN_NAME`,
				`COLUMN_TYPE`,
				`COLUMN_DEFAULT`,
				CASE`IS_NULLABLE` WHEN 'Yes' THEN '是' ELSE '否'END IS_NULLABLE,
				CASE `EXTRA` WHEN 'auto_increment' THEN '是' ELSE ' ' END EXTRA,
				`COLUMN_COMMENT`
			FROM
				information_schema.`COLUMNS`
			WHERE
				TABLE_NAME = '" . $currentTable . "'
			AND TABLE_SCHEMA = '" . $selfConf['db'] . "'"; //查询每一张表的字段名，数据类型，默认值，是否允许为空，是否递增以及备注
    $ret_column = $dbone->query($sql_column);
    if (!$ret_column) {
        echo '未查询到字段名等数据';
        die;
    }
    ?>
    <table class="gridtable">
        <tr>
            <th>字段序号</th>
            <th>字段名</th>
            <th>数据类型</th>
            <th>默认值</th>
            <th>允许非空</th>
            <th>是否递增</th>
            <th>备注</th>
        </tr>
        <?php
        $m = 1;
        while ($row_colunm = mysqli_fetch_assoc($ret_column)) {
            ?>
            <tr>
                <td><?php echo $m; ?></td>
                <td><?php echo $row_colunm['COLUMN_NAME'] ?></td>
                <td><?php echo $row_colunm['COLUMN_TYPE'] ?></td>
                <td><?php echo $row_colunm['COLUMN_DEFAULT'] ?></td>
                <td><?php echo $row_colunm['IS_NULLABLE'] ?></td>
                <td><?php echo $row_colunm['EXTRA'] ?></td>
                <td><?php echo $row_colunm['COLUMN_COMMENT'] ?></td>
            </tr>
            <?php
            $m++;
        }
        ?>
    </table>
    </div>
    </div>
    <?php
    $i++;
}


$dbtwo = $result['dbtwo']['conn'];
$dbtwo->query('set names utf8');
$sql_tables = "SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.`TABLES` WHERE TABLE_SCHEMA = '" . $sourceConf['db'] . "'"; //查询表名及其备注
$ret_tables = $dbtwo->query($sql_tables);
if (!$ret_tables) {
    echo '未查询到表单';
    die;
}
?>

</div>

<div class="col-lg-6 col-md-6 col-sm-6 col-xs-6"  >
	<div class="page-header">
      <h1><?php  echo $sourceConf['host']."  ".$sourceConf['db'] ; ?></h1>
    </div>
    
<?php

$i = 1;
while ($row_tables = mysqli_fetch_assoc($ret_tables)) {
    $currentTable = $row_tables['TABLE_NAME'];// 表的名称
?>
<div class="panel panel-info">
      <div class="panel-heading">
        <h3 class="panel-title">表序号：<?php echo $i . PHP_EOL . "&nbsp;&nbsp;"; ?> 
        <b>表名：<?php echo $row_tables['TABLE_NAME'] ?></b>　　
        <b>备注：<?php echo $row_tables['TABLE_COMMENT'] ?></b>
        </h3>
      </div>
      <div class="panel-body">
    <?php
    $sql_column = "SELECT
				`COLUMN_NAME`,
				`COLUMN_TYPE`,
				`COLUMN_DEFAULT`,
				CASE`IS_NULLABLE` WHEN 'Yes' THEN '是' ELSE '否'END IS_NULLABLE,
				CASE `EXTRA` WHEN 'auto_increment' THEN '是' ELSE ' ' END EXTRA,
				`COLUMN_COMMENT`
			FROM
				information_schema.`COLUMNS`
			WHERE
				TABLE_NAME = '" . $currentTable . "'
			AND TABLE_SCHEMA = '" . $sourceConf['db'] . "'"; //查询每一张表的字段名，数据类型，默认值，是否允许为空，是否递增以及备注
    $ret_column = $dbtwo->query($sql_column);
    if (!$ret_column) {
        echo '未查询到字段名等数据';
        die;
    }
    ?>
    <table class="gridtable">
        <tr>
            <th>字段序号</th>
            <th>字段名</th>
            <th>数据类型</th>
            <th>默认值</th>
            <th>允许非空</th>
            <th>是否递增</th>
            <th>备注</th>
        </tr>
        <?php
        $m = 1;
        while ($row_colunm = mysqli_fetch_assoc($ret_column)) {
            ?>
            <tr>
                <td><?php echo $m; ?></td>
                <td><?php echo $row_colunm['COLUMN_NAME'] ?></td>
                <td><?php echo $row_colunm['COLUMN_TYPE'] ?></td>
                <td><?php echo $row_colunm['COLUMN_DEFAULT'] ?></td>
                <td><?php echo $row_colunm['IS_NULLABLE'] ?></td>
                <td><?php echo $row_colunm['EXTRA'] ?></td>
                <td><?php echo $row_colunm['COLUMN_COMMENT'] ?></td>
            </tr>
            <?php
            $m++;
        }
        ?>
    </table>
    </div>
    </div>
    <?php
    $i++;
    
    
}
?>
</div>



<div class="btn btn-info" style="text-align: center; margin: 30px 0; "><a href="#" >返回顶部</a></div>

<div style="color: #ff004e;">
    <div class="">
        服务器IP地址 <?php echo $_SERVER['SERVER_ADDR'] . "\r\n\n" . "<br>"; ?>

        服务器域名 <?php echo $_SERVER['SERVER_NAME'] . "<br>"; ?>

        服务器端口 <?php echo $_SERVER['SERVER_PORT'] . "<br>"; ?>

        服务器版本 <?php echo php_uname('s') . php_uname('r') . "<br>"; ?>

        服务器操作系统 <?php echo php_uname() . "<br>"; ?>

        PHP版本 <?php echo PHP_VERSION . "<br>"; ?>

        获取PHP安装路径： <?php echo DEFAULT_INCLUDE_PATH . "<br>"; ?>

        获取当前文件绝对路径： <?php echo __FILE__ . "<br>"; ?>

        获取Http请求中Host值： <?php echo $_SERVER["HTTP_HOST"] . "<br>"; ?>

        获取Zend版本： <?php echo Zend_Version() . "<br>"; ?>

        <!--        Laravel版本 --><?php //echo $laravel = app();
        //        $laravel::VERSION; ?>

        PHP运行方式 <?php echo php_sapi_name() . "<br>"; ?>

        服务器当前时间 <?php echo date("Y-m-d H:i:s") . "<br>"; ?>

        最大上传限制 <?php echo get_cfg_var("upload_max_filesize") ? get_cfg_var("upload_max_filesize") : "不允许" . "<br>"; ?>

        最大执行时间 <?php echo get_cfg_var("max_execution_time") . "秒 " . "<br>"; ?>

        脚本运行占用最大内存 <?php echo get_cfg_var("memory_limit") ? get_cfg_var("memory_limit") : "无" . "<br>"; ?>

        获取服务器解译引擎： <?php echo $_SERVER['SERVER_SOFTWARE'] . "<br>"; ?>

        获取服务器CPU数量： <?php echo $_SERVER['PROCESSOR_IDENTIFIER'] . "<br>"; ?>

        获取服务器系统目录： <?php echo $_SERVER['SYSTEMROOT'] . "<br>"; ?>

        获取服务器域名（主机名）：<?php echo $_SERVER['SERVER_NAME'] . "(建议使用：" . $_SERVER["HTTP_HOST"] . ")" . "<br>"; ?>

        获取用户域名： <?php echo $_SERVER['USERDOMAIN'] . "<br>"; ?>

        获取服务器语言： <?php echo $_SERVER['HTTP_ACCEPT_LANGUAGE'] . "<br>"; ?>

        获取服务器Web端口： <?php echo $_SERVER['SERVER_PORT'] . "<br>"; ?>

        获取请求页面时通信协议的名称和版本： <?php echo $_SERVER['SERVER_PROTOCOL'] . "<br>"; ?>
    </div>
</div>
 </div>

    
</body>
</html>

