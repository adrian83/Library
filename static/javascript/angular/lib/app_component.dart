import 'package:angular2/core.dart';

class Author {
	String firstName = "Henryk";
	String lastName = "Sienkiewicz";
}


@Component(
    selector: 'my-app',
    template: '''
    	<h1>My First Angular 2 App</h1><div>{{author.firstName}} {{author.lastName}}</div>
	    <div>
	    	<input [(ngModel)]="author.firstName"/>
	    	<br/>
	    	<input [(ngModel)]="author.lastName"/>
	    	<br/>
	    	<input type="submit" />
	    </div>
    '''
    )
class AppComponent {
	Author author = new Author();
}
