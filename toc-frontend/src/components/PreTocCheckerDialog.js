import {
    MDBModal,
    MDBModalDialog,
    MDBModalContent,
    MDBModalHeader,
    MDBModalTitle,
    MDBModalBody,
    MDBModalFooter,
    MDBBtn
} from 'mdb-react-ui-kit';

export default function PreTocCheckerDialog({ 
    showPreTocCheckerDialog,
    closePreTocChecker,
    actionText
    }) {

    //const [basicModal, setBasicModal] = useState(false);
    //const toggleOpen = () => setBasicModal(!basicModal);

    return (
        <MDBModal staticBackdrop show={showPreTocCheckerDialog} tabIndex='-1'>
            <MDBModalDialog size="lg">
                <MDBModalContent>
                    <MDBModalHeader>
                        <MDBModalTitle>Pre-TOC Checker</MDBModalTitle>
                        <MDBBtn className='btn-close' color='none' onClick={closePreTocChecker}></MDBBtn>
                    </MDBModalHeader>
                    <MDBModalBody>
                        <div className='container-fluid'>
                            <div className='row'>
                                <div className='col-md-8'>File</div>
                                <div className='col-md-4'>...</div>
                            </div>
                            <div className='row'>
                                <div className='col-md-8'>Web</div>
                                <div className='col-md-4'>...</div>
                            </div>
                            <div className='row'>
                                <div className='col-md-8'>Mail</div>
                                <div className='col-md-4'>...</div>
                            </div>
                            <div className='row'>
                                <div className='col-md-8'>Database</div>
                                <div className='col-md-4'>...</div>
                            </div>
                        </div>
                    </MDBModalBody>
                    <MDBModalFooter>
                        <MDBBtn onClick={closePreTocChecker} color="secondary">Exit</MDBBtn>
                        <MDBBtn onClick={closePreTocChecker}>{actionText}</MDBBtn>
                    </MDBModalFooter>
                </MDBModalContent>
            </MDBModalDialog>
        </MDBModal>
    )
}