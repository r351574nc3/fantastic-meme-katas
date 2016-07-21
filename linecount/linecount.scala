import scala.collection.mutable.{Map}
import scala.io.{Source}
import java.io.{File}
import java.util.{HashMap}

object linecount {
    
    def run(path:String) = {
        implicit val codec = scalax.io.Codec.UTF8
    	var frequencies:Map[String, Int] = Map()
        for (line <- Source.fromFile(path).getLines) {
            var count = frequencies.get(line).getOrElse(0) + 1
            frequencies.put(line, count)
        }

        val keys = frequencies.keys
        val sorted_keys = keys.sortWith((a, b) => 
        		frequencies(a) < frequencies(b)              
            )

        var newpath = path + ".sorted"
        
        var p = new PrintWriter(newpath)
        try {
             
        }
        catch (Exception e) {
        }
        finally {
        }
	    Some(new PrintWriter(newpath)).foreach{
	        p => { 
	        	sorted_keys.foreach {
	    			p.write(frequencies.get(_).get + " " + _); p.close}
        	newpath.write()
        }
        println("Keys: " + keys)
    } 

    def main(args: Array[String]) {
        if (args.length  < 1) {
        	println("A filename is required")
        	return
        }

        val path = args(0)
        if (path == null || path.isEmpty()) {
        	println("Please give a valid filename")
        }
    	run(path)
    }
}
